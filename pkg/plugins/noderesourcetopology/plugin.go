package noderesourcetopology

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"

	topologyclientset "github.com/gocrane/api/pkg/generated/clientset/versioned"
	informers "github.com/gocrane/api/pkg/generated/informers/externalversions"
	listerv1alpha1 "github.com/gocrane/api/pkg/generated/listers/topology/v1alpha1"

	"github.com/mokaz111/mokazSche/pkg/plugins/apis/config"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"time"
)

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "NodeResourceTopologyMatch"

	// stateKey is the key in CycleState to NodeResourcesTopology.
	stateKey framework.StateKey = Name
)

// TopologyMatch plugin which run simplified version of TopologyManager's admit handler
type TopologyMatch struct {
	PodTopologyCache
	handle                 framework.Handle
	lister                 listerv1alpha1.NodeResourceTopologyLister
	topologyAwareResources sets.String
}

// New initializes a new plugin and returns it.
func New(args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	klog.V(2).InfoS("Creating new TopologyMatch plugin")
	cfg, ok := args.(*config.NodeResourceTopologyMatchArgs)
	if !ok {
		return nil, fmt.Errorf("want args to be of type NodeResourceTopologyMatchArgs, got %T", args)
	}

	ctx := context.TODO()
	client, err := topologyclientset.NewForConfig(handle.KubeConfig())
	if err != nil {
		klog.ErrorS(err, "Failed to create clientSet for NodeTopologyResource", "kubeConfig", handle.KubeConfig())
		return nil, err
	}

	lister, err := initTopologyInformer(ctx, client)
	if err != nil {
		return nil, err
	}

	topologyMatch := &TopologyMatch{
		PodTopologyCache:       NewPodTopologyCache(ctx, 30*time.Minute),
		handle:                 handle,
		lister:                 lister,
		topologyAwareResources: sets.NewString(cfg.TopologyAwareResources...),
	}

	return topologyMatch, nil
}

// Name returns name of the plugin. It is used in logs, etc.
func (tm *TopologyMatch) Name() string {
	return Name
}

func initTopologyInformer(
	ctx context.Context,
	client topologyclientset.Interface,
) (listerv1alpha1.NodeResourceTopologyLister, error) {
	topologyInformerFactory := informers.NewSharedInformerFactory(client, 0)
	nrtlister := topologyInformerFactory.Topology().V1alpha1().NodeResourceTopologies().Lister()

	klog.V(4).InfoS("Start nodetopologyInformer")
	topologyInformerFactory.Start(ctx.Done())
	topologyInformerFactory.WaitForCacheSync(ctx.Done())
	return nrtlister, nil
}
