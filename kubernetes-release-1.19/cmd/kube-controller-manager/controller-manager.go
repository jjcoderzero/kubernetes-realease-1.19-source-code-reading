// 控制器管理器负责监视副本控制器，并创建相应的pods以实现所需的状态。它使用API侦听新的控制器并创建/删除pods。
package main

import (
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // 加载所有prometheus client-go 插件
	_ "k8s.io/component-base/metrics/prometheus/version"  // 版本信息注册
	"k8s.io/kubernetes/cmd/kube-controller-manager/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewControllerManagerCommand()

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
