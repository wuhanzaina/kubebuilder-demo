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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of App. Edit app_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

/*
是 Kubebuilder 标记（markers），用于 controller-gen 工具自动生成代码和 CRD 文件。
这些“注解”不会影响运行时行为，但会影响生成的 Kubernetes CRD YAML 和 Go 类型

执行 make manifests Kubebuilder 会读取这些注解，然后生成如下内容：

# crds/apps.wzy.com_apps.yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: apps.apps.wzy.com
spec:
  ...
  subresources:
    status: {}  # <-- 来自 +kubebuilder:subresource:status
*/
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// App is the Schema for the apps API
// 定义 crd的对象
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`   // crd 核心内容
	Status AppStatus `json:"status,omitempty"` //crd现阶段的状态
}

//+kubebuilder:object:root=true

// AppList contains a list of App
// 定义 crd的列表
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
