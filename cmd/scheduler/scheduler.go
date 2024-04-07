package main

import (
	"fmt"
	"github.com/mokaz111/mokazSche/pkg/plugins/dynamic"
	"github.com/mokaz111/mokazSche/pkg/plugins/noderesourcetopology"
	"k8s.io/component-base/logs"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cmd := app.NewSchedulerCommand(
		app.WithPlugin(dynamic.Name, dynamic.NewDynamicScheduler),
		app.WithPlugin(noderesourcetopology.Name, noderesourcetopology.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
