// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package migrate // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"

import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/alias"

// SignalType allows for type constraints in order
// to apply to potential type defined strings.
type SignalType interface {
	~string
}

// SignalNameChange allows for migrating types that
// implement the `alias.NamedSignal` interface.
type SignalNameChange struct {
	updates  map[string]string
	rollback map[string]string
}

type SignalNameChangeSlice []*SignalNameChange

// NewSignalNameChange will create a `Signal` that will check the provided mappings if it can update a `alias.NamedSignal`
// and if no values are provided for `matches`, then all values will be updated.
func NewSignalNameChange[Key SignalType, Value SignalType](mappings map[Key]Value) *SignalNameChange {
	sig := &SignalNameChange{
		updates:  make(map[string]string, len(mappings)),
		rollback: make(map[string]string, len(mappings)),
	}
	for k, v := range mappings {
		sig.updates[string(k)] = string(v)
		sig.rollback[string(v)] = string(k)
	}
	return sig
}

func (s SignalNameChange) IsMigrator() {}


func (s *SignalNameChange) Apply(signal alias.NamedSignal) {
	s.Do(StateSelectorApply, signal)
}

func (s *SignalNameChange) Rollback(signal alias.NamedSignal) {
	s.Do(StateSelectorRollback, signal)
}

func (s *SignalNameChange) Do(ss StateSelector, signal alias.NamedSignal) {
	var (
		name    string
		matched bool
	)
	switch ss {
	case StateSelectorApply:
		name, matched = s.updates[signal.Name()]
	case StateSelectorRollback:
		name, matched = s.rollback[signal.Name()]
	}
	if matched {
		signal.SetName(name)
	}
}

func NewSignalNameChangeSlice(changes ...*SignalNameChange) *SignalNameChangeSlice {
	values := new(SignalNameChangeSlice)
	for _, c := range changes {
		(*values) = append((*values), c)
	}
	return values
}

// todo make this return an err
func (slice *SignalNameChangeSlice) Apply(signal alias.NamedSignal) {
	slice.Do(StateSelectorApply, signal)
}

func (slice *SignalNameChangeSlice) Rollback(signal alias.NamedSignal) {
	slice.Do(StateSelectorRollback, signal)
}

func (slice *SignalNameChangeSlice) Do(ss StateSelector, signal alias.NamedSignal) {
	for i := 0; i < len((*slice)); i++ {
		switch ss {
		case StateSelectorApply:
			(*slice)[i].Apply(signal)
		case StateSelectorRollback:
			(*slice)[len((*slice))-i-1].Rollback(signal)
		}
	}
}
