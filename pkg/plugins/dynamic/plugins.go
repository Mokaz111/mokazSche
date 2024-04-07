package dynamic

import (
	"fmt"
	"github.com/mokaz111/mokazSche/pkg/plugins/apis/config"
	"github.com/mokaz111/mokazSche/pkg/plugins/apis/policy"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "Dynamic"
)

// Dynamic-scheduler is a real load-aware scheduler plugin.
type DynamicScheduler struct {
	handle          framework.Handle
	schedulerPolicy *policy.DynamicSchedulerPolicy
}

// Name returns name of the plugin.
func (ds *DynamicScheduler) Name() string {
	return Name
}

// NewDynamicScheduler returns a mokaz Scheduler object.
func NewDynamicScheduler(plArgs runtime.Object, h framework.Handle) (framework.Plugin, error) {
	args, ok := plArgs.(*config.DynamicArgs)
	if !ok {
		return nil, fmt.Errorf("want args to be of type DynamicArgs, got %T.", plArgs)
	}

	schedulerPolicy, err := LoadPolicyFromFile(args.PolicyConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler policy from config file: %v", err)
	}

	return &DynamicScheduler{
		schedulerPolicy: schedulerPolicy,
		handle:          h,
	}, nil
}
