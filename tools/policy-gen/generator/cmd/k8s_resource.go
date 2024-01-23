package cmd

import (
	"os"
	"path/filepath"
	"text/template"
)

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/apache/dubbo-kubernetes/tools/policy-gen/generator/pkg/parse"
	"github.com/apache/dubbo-kubernetes/tools/policy-gen/generator/pkg/save"
)

func newK8sResource(rootArgs *args) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "k8s-resource",
		Short: "Generate a k8s model resource for the policy",
		Long:  "Generate a k8s model resource for the policy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			policyName := filepath.Base(rootArgs.pluginDir)
			policyPath := filepath.Join(rootArgs.pluginDir, "api", rootArgs.version, policyName+".go")
			if _, err := os.Stat(policyPath); err != nil {
				return err
			}

			pconfig, err := parse.Policy(policyPath)
			if err != nil {
				return err
			}

			pconfig.GoModule = rootArgs.goModule

			k8sPath := filepath.Join(rootArgs.pluginDir, "k8s", rootArgs.version)
			if err := os.MkdirAll(k8sPath, 0o755); err != nil {
				return err
			}

			k8sTypesPath := filepath.Join(k8sPath, "zz_generated.types.go")
			if err := save.GoTemplate(customResourceTemplate, pconfig, k8sTypesPath); err != nil {
				return err
			}

			gvInfoPath := filepath.Join(k8sPath, "groupversion_info.go")
			if err := save.GoTemplate(groupVersionInfoTemplate, pconfig, gvInfoPath); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

// customResourceTemplate for creating a Kubernetes CRD to wrap a Dubbo resource.
var customResourceTemplate = template.Must(template.New("custom-resource").Parse(`
// Generated by tools/policy-gen
// Run "make generate" to update this file.

{{ $tk := "` + "`" + `" }}

// nolint:whitespace
package {{.Package}}

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	core_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	policy "{{.GoModule}}/pkg/plugins/policies/{{.NameLower}}/api/{{.Package}}"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
	{{- if not .SkipRegistration }}
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/registry"
	{{- end }}
	"github.com/apache/dubbo-kubernetes/pkg/plugins/runtime/k8s/metadata"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
type {{.Name}} struct {
	metav1.TypeMeta   {{ $tk }}json:",inline"{{ $tk }}
	metav1.ObjectMeta {{ $tk }}json:"metadata,omitempty"{{ $tk }}

	// Spec is the specification of the Dubbo {{ .Name }} resource.
    // +kubebuilder:validation:Optional
	Spec   *policy.{{.Name}} {{ $tk }}json:"spec,omitempty"{{ $tk }}
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type {{.Name}}List struct {
	metav1.TypeMeta {{ $tk }}json:",inline"{{ $tk }}
	metav1.ListMeta {{ $tk }}json:"metadata,omitempty"{{ $tk }}
	Items           []{{.Name}} {{ $tk }}json:"items"{{ $tk }}
}

func (cb *{{.Name}}) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *{{.Name}}) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *{{.Name}}) GetMesh() string {
	if mesh, ok := cb.ObjectMeta.Labels[metadata.DubboMeshLabel]; ok {
		return mesh
	} else {
		return core_model.DefaultMesh
	}
}

func (cb *{{.Name}}) SetMesh(mesh string) {
	if cb.ObjectMeta.Labels == nil {
		cb.ObjectMeta.Labels = map[string]string{}
	}
	cb.ObjectMeta.Labels[metadata.DubboMeshLabel] = mesh
}

func (cb *{{.Name}}) GetSpec() (core_model.ResourceSpec, error) {
	return cb.Spec, nil
}

func (cb *{{.Name}}) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
		cb.Spec = nil
		return
	}

	if _, ok := spec.(*policy.{{.Name}}); !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

	cb.Spec = spec.(*policy.{{.Name}})
}

func (cb *{{.Name}}) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *{{.Name}}List) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

{{if not .SkipRegistration}}
func init() {
	SchemeBuilder.Register(&{{.Name}}{}, &{{.Name}}List{})
	registry.RegisterObjectType(&policy.{{.Name}}{}, &{{.Name}}{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "{{.Name}}",
		},
	})
	registry.RegisterListType(&policy.{{.Name}}{}, &{{.Name}}List{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "{{.Name}}List",
		},
	})
}
{{- end }} {{/* .SkipRegistration */}}
`))

var groupVersionInfoTemplate = template.Must(template.New("groupversion-info").Parse(`
// Package {{.Package}} contains API Schema definitions for the mesh {{.Package}} API group
// +groupName=dubbo.io
package {{.Package}}

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "dubbo.io", Version: "{{.Package}}"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
`))
