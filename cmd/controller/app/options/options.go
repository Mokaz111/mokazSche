package options

import (
	"fmt"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/config/options"
	controllerappconfig "mokazSche/cmd/controller/app/config"
	annotatorconfig "mokazSche/pkg/controller/annotator/config"
	"mokazSche/pkg/controller/prometheus"
	dynamicscheduler "mokazSche/pkg/plugins/dynamic"
	"mokazSche/pkg/utils"
	"time"
)

const (
	ControllerUserAgent = "mokaz-scheduler-controller"
)

type Options struct {
	*annotatorconfig.AnnotatorConfiguration
	LeaderElection *componentbaseconfig.LeaderElectionConfiguration
	master         string
	kubeconfig     string
	healthPort     string
}

func (o *Options) Config() (*controllerappconfig.Config, error) {
	var kubeconfig *rest.Config
	var err error
	if err := o.Validate(); err != nil {
		return nil, err
	}
	c := &controllerappconfig.Config{}
	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	c.Policy, err = dynamicscheduler.LoadPolicyFromFile(o.PolicyConfigPath)
	if err != nil {
		return nil, err
	}

	if o.kubeconfig == "" {
		kubeconfig, err = rest.InClusterConfig()
	} else {
		// Build config from configfile
		kubeconfig, err = clientcmd.BuildConfigFromFlags(o.master, o.kubeconfig)
	}
	if err != nil {
		return nil, err
	}

	c.KubeClient, err = clientset.NewForConfig(rest.AddUserAgent(kubeconfig, ControllerUserAgent))
	if err != nil {
		return nil, err
	}

	c.LeaderElectionClient = clientset.NewForConfigOrDie(rest.AddUserAgent(kubeconfig, "leader-election"))

	c.PromClient, err = prometheus.NewPromClient(o.PrometheusAddr)
	if err != nil {
		return nil, err
	}
	c.KubeInformerFactory = NewInformerFactory(c.KubeClient, 0)

	c.HealthPort = o.healthPort
	return c, nil
}

func (o *Options) Validate() error {
	return nil
}

// ApplyTo fills up Annotator config with options.
func (o *Options) ApplyTo(c *controllerappconfig.Config) error {
	c.AnnotatorConfig = o.AnnotatorConfiguration
	c.LeaderElection = o.LeaderElection
	return nil
}

func (o *Options) Flags(flag *pflag.FlagSet) error {
	if flag == nil {
		return fmt.Errorf("nil pointer")
	}

	flag.StringVar(&o.PolicyConfigPath, "policy-config-path", o.PolicyConfigPath, "Path to annotator policy config")
	flag.StringVar(&o.PrometheusAddr, "prometheus-address", o.PrometheusAddr, "The address of prometheus, from which we can pull metrics data.")
	flag.Int32Var(&o.BindingHeapSize, "binding-heap-size", o.BindingHeapSize, "Max size of binding heap size, used to store hot value data.")
	flag.Int32Var(&o.ConcurrentSyncs, "concurrent-syncs", o.ConcurrentSyncs, "The number of annotator controller workers that are allowed to sync concurrently.")
	flag.StringVar(&o.kubeconfig, "kubeconfig", o.kubeconfig, "Path to kubeconfig file with authorization information")
	flag.StringVar(&o.master, "master", o.master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	flag.StringVar(&o.healthPort, "health-port", o.healthPort, "The port of health check")

	options.BindLeaderElectionFlags(o.LeaderElection, flag)
	return nil
}
func NewOptions() (*Options, error) {
	o := &Options{
		AnnotatorConfiguration: &annotatorconfig.AnnotatorConfiguration{
			BindingHeapSize:  1024,
			ConcurrentSyncs:  1,
			PolicyConfigPath: "/etc/kubernetes/policy.yaml",
		},
		LeaderElection: &componentbaseconfig.LeaderElectionConfiguration{
			LeaderElect:       true,
			LeaseDuration:     metav1.Duration{Duration: 15 * time.Second},
			RenewDeadline:     metav1.Duration{Duration: 10 * time.Second},
			RetryPeriod:       metav1.Duration{Duration: 2 * time.Second},
			ResourceLock:      "leases",
			ResourceName:      "mokaz-scheduler-conreoller",
			ResourceNamespace: utils.GetSystemNamespace(),
		},
		healthPort: "8090",
	}
	return o, nil
}
