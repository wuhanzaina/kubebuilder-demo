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

// Package v1 contains API Schema definitions for the apps v1 API group
// +kubebuilder:object:generate=true
// +groupName=apps.wzy.com
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

/*
Kubebuilder 构建 Kubernetes 自定义资源（CRD）控制器的 Go 代码文件，位于 API 包中（如 api/v1/groupversion_info.go）。
它的目的是定义并注册一个 API 组（Group）和版本（Version），以便让控制器、Kubernetes API Server、客户端等都能识别自定义资源对象。
*/
var (
	/*
			定义了当前 API 的组（Group）和版本（Version），这是 Kubernetes 中识别资源的基本单位。例如：
		apiVersion: apps.wzy.com/v1
		kind: App
			解释：
			Group: 是 API 分组名（可自定义，比如 apps.wzy.com）。
			Version: 是版本号（比如 v1）。
			GroupVersion: 是这两个值的组合，后续注册资源类型和反序列化等都基于这个组合。
	*/
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "apps.wzy.com", Version: "v1"}
	/*
	   创建一个指向 scheme.Builder 的指针，用于后续向 K8s 的 runtime.Scheme 中注册自定义资源类型（CRD）。
	*/
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	/*
	   这是一个快捷方式，将 SchemeBuilder.AddToScheme 方法暴露出来，方便在其他地方调用。
	*/
	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
