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

/*
这个文件定义了一个控制器 AppReconciler，它的职责是：

只要 App 资源有任何变化（创建、更新、删除），就自动执行 Reconcile() 方法，确保 当前集群状态 与 期望状态（用户定义的 App YAML）一致。
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "com.wzy.onedemo/api/v1"
)

// AppReconciler reconciles a App object
// 这个结构体是一个 控制器，你要在 main.go 中注册它。
type AppReconciler struct {
	// 封装了对 K8s API 的访问能力，比如 get/list/update/delete 对象
	client.Client
	// 用于转换对象的 GVK 与 Go 类型
	//结构体中的指针变量字段 *xxx 结构体
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.wzy.com,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.wzy.com,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.wzy.com,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
/*
这个是每次资源变化时被自动调用的“对比与同步”逻辑：

	ctx: 上下文，可用于日志、取消等

	req: 包含本次触发的对象的 NamespacedName（即 namespace/name）

你会在这里写入实际的业务逻辑，比如：

	获取 App 对象

	检查其字段

	创建/更新 Deployment、Service 等 K8s 原生资源来“实现” App 的期望状态

	更新 App 的 .status 字段反映结果
*/
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
// controller 注册到 manager中
// 告诉控制器,我关注的资源是 App，请监听它的变化，变化后就自动调用 Reconcile() 方法。
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&appsv1.App{}).Complete(r)
}
