package k8s

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
)

type nodeNameProvider interface {
	NodeName(context.Context) (string, error)
}

var _ nodeNameProvider = (*nodeNameProviderImpl)(nil)

type nodeNameProviderImpl struct {
	logger *zap.Logger
	client k8s.Interface
}

func (p *nodeNameProviderImpl) namespace() string {
	namespacePath := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	val, err := ioutil.ReadFile(namespacePath)
	if err == nil && val != nil {
		return string(val)
	}
	p.logger.Warn("Could not fetch k8s namespace, using 'default'", zap.Error(err))
	return "default"
}

func (p *nodeNameProviderImpl) NodeName(ctx context.Context) (string, error) {
	namespace := p.namespace()
	podName, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("could not fetch our hostname: %w", err)
	}

	pod, err := p.client.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return pod.Spec.NodeName, nil
}

func newNodeNameProvider(logger *zap.Logger) (nodeNameProvider, error) {
	client, err := k8sconfig.MakeClient(k8sconfig.APIConfig{
		AuthType: k8sconfig.AuthTypeServiceAccount,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to build k8s client: %w", err)
	}

	return &nodeNameProviderImpl{
		client: client,
	}, nil
}
