package config

import (
	"github.com/mokaz111/mokazSche/pkg/plugins/apis/policy"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"

	annotatorconfig "github.com/mokaz111/mokazSche/pkg/controller/annotator/config"
	componentbaseconfig "k8s.io/component-base/config"

	prom "github.com/mokaz111/mokazSche/pkg/controller/prometheus"
)

type Config struct {
	AnnotatorConfig *annotatorconfig.AnnotatorConfiguration
	LeaderElection  *componentbaseconfig.LeaderElectionConfiguration
	// KubeInformerFactory gives access to kubernetes informers for the controller.
	KubeInformerFactory informers.SharedInformerFactory
	// KubeClient is the general kube client.
	KubeClient clientset.Interface
	// PromClient is used for getting metric data from Prometheus.
	PromClient prom.PromClient
	// Policy is a collection of scheduler policies.
	Policy *policy.DynamicSchedulerPolicy
	// EventRecorder is the event sink
	EventRecorder record.EventRecorder
	// LeaderElectionClient is the client used for leader election
	LeaderElectionClient *clientset.Clientset
	// HealthPort is server port used for health check
	HealthPort string
}
type completedConfig struct {
	*Config
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	cc := completedConfig{c}

	return &CompletedConfig{&cc}
}
