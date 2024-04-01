package main

import (
	goflag "flag"
	"fmt"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"math/rand"
	"mokazSche/cmd/controller/app"
	"os"
	"time"

	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cmd := app.NewControllerCommand()

	klog.InitFlags(nil)

	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
