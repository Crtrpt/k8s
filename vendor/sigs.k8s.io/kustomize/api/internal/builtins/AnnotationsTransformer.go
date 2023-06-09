// Code generated by pluginator on AnnotationsTransformer; DO NOT EDIT.
// pluginator {(devel)  unknown   }

package builtins

import (
	"sigs.k8s.io/kustomize/api/filters/annotations"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

// Add the given annotations to the given field specifications.
type AnnotationsTransformerPlugin struct {
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *AnnotationsTransformerPlugin) Config(
	_ *resmap.PluginHelpers, c []byte) (err error) {
	p.Annotations = nil
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *AnnotationsTransformerPlugin) Transform(m resmap.ResMap) error {
	if len(p.Annotations) == 0 {
		return nil
	}
	return m.ApplyFilter(annotations.Filter{
		Annotations: p.Annotations,
		FsSlice:     p.FieldSpecs,
	})
}

func NewAnnotationsTransformerPlugin() resmap.TransformerPlugin {
	return &AnnotationsTransformerPlugin{}
}
