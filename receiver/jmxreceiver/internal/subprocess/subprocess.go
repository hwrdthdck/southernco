// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subprocess

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

const (
	defaultRestartDelay    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
	noPid                  = -1
)

// Config exported to be used by jmx metric receiver.
type Config struct {
	ExecutablePath       string            `mapstructure:"executable_path"`
	Args                 []string          `mapstructure:"args"`
	EnvironmentVariables map[string]string `mapstructure:"environment_variables"`
	StdInContents        string            `mapstructure:"stdin_contents"`
	RestartOnError       bool              `mapstructure:"restart_on_error"`
	RestartDelay         *time.Duration    `mapstructure:"restart_delay"`
	ShutdownTimeout      *time.Duration    `mapstructure:"shutdown_timeout"`
}

// Subprocess exported to be used by jmx metric receiver.
type Subprocess struct {
	stdout         chan string
	cancel         context.CancelFunc
	config         *Config
	envVars        []string
	logger         *zap.Logger
	pid            pid
	shutdownSignal chan struct{}
	// configurable for testing purposes
	sendToStdIn func(string, io.Writer) error
}

type pid struct {
	pid     int
	pidLock sync.Mutex
}

func (p *pid) setPid(pid int) {
	p.pidLock.Lock()
	defer p.pidLock.Unlock()
	p.pid = pid
}

func (p *pid) getPid() int {
	p.pidLock.Lock()
	defer p.pidLock.Unlock()
	return p.pid
}

func (subprocess *Subprocess) Pid() int {
	pid := subprocess.pid.getPid()
	if pid == 0 {
		return noPid
	}
	return pid
}

// NewSubprocess exported to be used by jmx metric receiver.
func NewSubprocess(conf *Config, logger *zap.Logger) *Subprocess {
	if conf.RestartDelay == nil {
		restartDelay := defaultRestartDelay
		conf.RestartDelay = &restartDelay
	}
	if conf.ShutdownTimeout == nil {
		shutdownTimeout := defaultShutdownTimeout
		conf.ShutdownTimeout = &shutdownTimeout
	}

	return &Subprocess{
		stdout:         make(chan string),
		pid:            pid{pid: noPid, pidLock: sync.Mutex{}},
		config:         conf,
		logger:         logger,
		shutdownSignal: make(chan struct{}),
		sendToStdIn:    sendToStdIn,
	}
}

const (
	starting     = "Starting"
	running      = "Running"
	shuttingDown = "ShuttingDown"
	stopped      = "Stopped"
	restarting   = "Restarting"
	errored      = "Errored"
)

func (subprocess *Subprocess) Stdout() chan string {
	return subprocess.stdout
}

func (subprocess *Subprocess) Start(ctx context.Context) error {
	var cancelCtx context.Context
	cancelCtx, subprocess.cancel = context.WithCancel(ctx)

	for k, v := range subprocess.config.EnvironmentVariables {
		joined := fmt.Sprintf("%v=%v", k, v)
		subprocess.envVars = append(subprocess.envVars, joined)
	}

	go func() {
		subprocess.run(cancelCtx) // will block for lifetime of process
		close(subprocess.shutdownSignal)
	}()
	return nil
}

// Shutdown is invoked during service shutdown.
func (subprocess *Subprocess) Shutdown(ctx context.Context) error {
	if subprocess.cancel == nil {
		return fmt.Errorf("no subprocess.cancel().  Has it been started properly?")
	}

	timeout := defaultShutdownTimeout
	if subprocess.config.ShutdownTimeout != nil {
		timeout = *subprocess.config.ShutdownTimeout
	}
	t := time.NewTimer(timeout)

	// Wait for the subprocess to exit or the timeout period to elapse
	select {
	case <-ctx.Done():
	case <-subprocess.shutdownSignal:
	case <-t.C:
		subprocess.logger.Warn("subprocess hasn't returned within shutdown timeout. May be zombied.",
			zap.String("timeout", fmt.Sprintf("%v", timeout)))
	}

	return nil
}

// A synchronization helper to ensure that signalWhenProcessReturned
// doesn't write to a closed channel
type processReturned struct {
	ReturnedChan chan error
	isOpen       *atomic.Bool
	lock         *sync.Mutex
}

func newProcessReturned() *processReturned {
	pr := processReturned{
		ReturnedChan: make(chan error),
		isOpen:       atomic.NewBool(true),
		lock:         &sync.Mutex{},
	}
	return &pr
}

func (pr *processReturned) signal(err error) {
	pr.lock.Lock()
	defer pr.lock.Unlock()
	if pr.isOpen.Load() {
		pr.ReturnedChan <- err
	}
}

func (pr *processReturned) close() {
	pr.lock.Lock()
	defer pr.lock.Unlock()
	if pr.isOpen.Load() {
		close(pr.ReturnedChan)
		pr.isOpen.Store(false)
	}
}

// Core event loop
func (subprocess *Subprocess) run(ctx context.Context) {
	var cmd *exec.Cmd
	var err error
	var stdin io.WriteCloser
	var stdout io.ReadCloser

	// writer is signalWhenProcessReturned() and closer is this loop, so we need synchronization
	processReturned := newProcessReturned()

	state := starting
	for {
		subprocess.logger.Debug("subprocess changed state", zap.String("state", state))

		switch state {
		case starting:
			cmd, stdin, stdout = createCommand(
				subprocess.config.ExecutablePath,
				subprocess.config.Args,
				subprocess.envVars,
			)

			go CollectStdout(bufio.NewScanner(stdout), subprocess.Stdout(), subprocess.logger)

			subprocess.logger.Debug("starting subprocess", zap.String("command", cmd.String()))
			err = cmd.Start()
			if err != nil {
				state = errored
				continue
			}
			subprocess.pid.setPid(cmd.Process.Pid)

			go signalWhenProcessReturned(cmd, processReturned)

			state = running
		case running:
			err = subprocess.sendToStdIn(subprocess.config.StdInContents, stdin)
			stdin.Close()
			if err != nil {
				state = errored
				continue
			}

			select {
			case err = <-processReturned.ReturnedChan:
				if err != nil && ctx.Err() == nil {
					err = fmt.Errorf("unexpected shutdown: %w", err)
					// We aren't supposed to shutdown yet so this is an error state.
					state = errored
					continue
				}
				// We must close this channel or can wait indefinitely at shuttingDown
				processReturned.close()
				state = shuttingDown
			case <-ctx.Done(): // context-based cancel.
				state = shuttingDown
			}
		case errored:
			subprocess.logger.Error("subprocess died", zap.Error(err))
			if subprocess.config.RestartOnError {
				subprocess.pid.setPid(-1)
				state = restarting
			} else {
				// We must close this channel or can wait indefinitely at shuttingDown
				processReturned.close()
				state = shuttingDown
			}
		case shuttingDown:
			if cmd.Process != nil {
				cmd.Process.Signal(syscall.SIGTERM)
			}
			<-processReturned.ReturnedChan
			stdout.Close()
			subprocess.pid.setPid(-1)
			state = stopped
		case restarting:
			stdout.Close()
			stdin.Close()
			time.Sleep(*subprocess.config.RestartDelay)
			state = starting
		case stopped:
			return
		}
	}
}

func signalWhenProcessReturned(cmd *exec.Cmd, pr *processReturned) {
	err := cmd.Wait()
	pr.signal(err)
}

func CollectStdout(stdoutScanner *bufio.Scanner, stdoutChan chan<- string, logger *zap.Logger) {
	for stdoutScanner.Scan() {
		text := stdoutScanner.Text()
		if text != "" {
			stdoutChan <- text
			logger.Debug(text)

			if statusCode := checkForStatusCode(text); statusCode != "" {
				logger.Error(text, zap.String("status_code", statusCode))
			}
		}
	}
	// Returns when stdout is closed when the process ends
}

func checkForStatusCode(text string) string {
	if strings.Contains(text, "Could not connect to remote JMX server") {
		return "UNAVAILABLE"
	} else if strings.Contains(text, "`interval` must be positive:") {
		return codes.InvalidArgument.String()
	} else if strings.Contains(text, "`otlp.timeout` must be positive:") {
		return codes.InvalidArgument.String()
	} else if strings.Contains(text, "jmx missing required fields:") {
		return codes.InvalidArgument.String()
	} else if strings.Contains(text, "failed to parse OTLPExporterConfig.Endpoint") {
		return codes.InvalidArgument.String()
	} else if strings.Contains(text, "failed to parse Endpoint") {
		return codes.InvalidArgument.String()
	} else if strings.Contains(text, "no subprocess.cancel().") {
		return codes.Internal.String()
	} else if strings.Contains(text, "subprocess hasn't returned within shutdown timeout.") {
		return codes.OutOfRange.String()
	} else if strings.Contains(text, "unexpected shutdown:") {
		return codes.Aborted.String()
	} else if strings.Contains(text, "Input pipe could not be created for subprocess") {
		return codes.Internal.String()
	} else if strings.Contains(text, "Output pipe could not be created for subprocess") {
		return codes.Internal.String()
	}
	return ""
}

func sendToStdIn(contents string, writer io.Writer) error {
	if contents == "" {
		return nil
	}

	_, err := writer.Write([]byte(contents))
	return err
}

func createCommand(execPath string, args, envVars []string) (*exec.Cmd, io.WriteCloser, io.ReadCloser) {
	cmd := exec.Command(execPath, args...)

	var env []string
	env = append(env, os.Environ()...)
	cmd.Env = append(env, envVars...)

	inReader, inWriter, err := os.Pipe()
	if err != nil {
		panic("Input pipe could not be created for subprocess")
	}

	cmd.Stdin = inReader

	outReader, outWriter, err := os.Pipe()
	if err != nil {
		panic("Output pipe could not be created for subprocess")
	}
	cmd.Stdout = outWriter
	cmd.Stderr = outWriter

	applyOSSpecificCmdModifications(cmd)

	return cmd, inWriter, outReader
}
