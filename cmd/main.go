/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	appsv1 "com.wzy.onedemo/api/v1"
	"com.wzy.onedemo/internal/controller"
	//+kubebuilder:scaffold:imports
)

var (
	/*
		对象 用来管理定义的g v k  和 go struct(crd) 映射 还有一些互相转换的方法
				什么是 scheme？
			scheme 是 Kubernetes 客户端（controller-runtime）用来进行 资源序列化 / 反序列化（encoding/decoding） 的注册表。

			它会告诉程序："我知道有一种叫 apps/v1.Deployment 的资源，它结构是这样的……"
			scheme 就像是一个资源类型的“字典”。
	*/
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {

	//将你的自定义资源（CRD）和 Kubernetes 标准资源注册到一个全局的 scheme 中。

	/*
		把 Kubernetes 标准资源（Pod、Service、Deployment 等）加入到 scheme 中。
		clientgoscheme 包含了所有内置类型（core/v1、apps/v1 等等）。
	*/
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	/*
		这一行 你可以添加自己的 CRD 类型或引用的资源类型。
		举个例子：
			如果你在 api/v1 中定义了一个 App 类型（Group: apps.my.domain，Version: v1，Kind: App）
	*/
	utilruntime.Must(appsv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

/*
┌──────────────────────────────┐
│  解析配置和命令行参数（flag）│
└────────────┬─────────────────┘

	↓

┌──────────────────────────────┐
│     创建 Manager（管理器）    │
└────────────┬─────────────────┘

	↓

┌──────────────────────────────┐
│ 注册 Reconciler 控制器（CRD） │
└────────────┬─────────────────┘

	↓

┌──────────────────────────────┐
│     配置健康检查和就绪检查    │
└────────────┬─────────────────┘

	↓

┌──────────────────────────────┐
│      启动 Manager 和监听器    │
│（阻塞主线程，直到接收中断） │
└──────────────────────────────┘
*/
func main() {
	//处理命令行参数
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	/*
		这三行是给程序注册命令行参数，支持用户用 --xxx=xxx 的方式传入值：
		./manager --metrics-bind-address=:9090 --leader-elect=true

	*/
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	/*
				初始化 zap 日志选项：

		Development: true 表示使用开发者友好的日志格式（比如带颜色、换行等）

		false 会更偏向 JSON 格式，适合生产日志收集系统（如 Loki）
	*/
	opts := zap.Options{
		Development: true,
	}
	//把 zap 的日志配置参数也注入到命令行中，让你可以传比如 --zap-log-level=debug 等参数。
	opts.BindFlags(flag.CommandLine)
	//解析命令行参数，把值赋给上面定义的变量，比如 metricsAddr
	flag.Parse()

	//设置日志
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	//完成命令行参数处理

	//创建manager
	/*
		ctrl.NewManager()创建一个 controller-runtime 的 Manager 实例，它负责生命周期管理、客户端缓存、事件监听、启动各控制器等。
		ctrl.GetConfigOrDie() 获取当前 Kubernetes 集群的配置信息（通过 ~/.kube/config 或 in-cluster 配置）。

		参数是一个 ctrl.Options 结构体，用于配置 Manager 的行为
		Scheme						所有注册到管理器的资源类型（比如 CRD 对象）
		MetricsBindAddress			metrics 暴露地址，供 Prometheus 等采集
		Port						用于 webhook（默认 9443），你可忽略此端口如果没启用 webhook
		HealthProbeBindAddress		健康探针和就绪探针监听的地址
		LeaderElection				是否启用 Leader 选举（多副本部署时只运行一个控制器）
		LeaderElectionID			选举锁 ID，通常设置为唯一值，如域名风格

	*/
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "a9ba9dc1.wzy.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}
	/*
		初始化并注册你自定义的controller 资源控制器，这里是 AppReconciler（用于处理 App CRD）。
		它实现了 Reconcile 接口（控制器核心逻辑），并调用 SetupWithManager 将它与 Manager 绑定。

	*/
	if err = (&controller.AppReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "App")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder
	/*
		健康检查与就绪检查
	*/
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	//启动程序
	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
