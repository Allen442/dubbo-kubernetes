package render

import (
	"fmt"
	"github.com/apache/dubbo-kubernetes/operator/cmd/validation"
	"github.com/apache/dubbo-kubernetes/operator/pkg/apis"
	"github.com/apache/dubbo-kubernetes/operator/pkg/component"
	"github.com/apache/dubbo-kubernetes/operator/pkg/manifest"
	"github.com/apache/dubbo-kubernetes/operator/pkg/util"
	"github.com/apache/dubbo-kubernetes/operator/pkg/util/clog"
	"github.com/apache/dubbo-kubernetes/operator/pkg/values"
	"io"
	"os"
	"strings"
)

func MergeInputs(filenames []string, flags []string) ([]values.Map, error) {
	ConfigBase, err := values.MapFromJSON([]byte(`{
	  "apiVersion": "install.dubbo.io/v1alpha1",
	  "kind": "DubboOperator",
	  "metadata": {},
	  "spec": {}
	}`))
	if err != nil {
		return nil, err
	}
	for i, fn := range filenames {
		var b []byte
		var err error
		if fn == "-" {
			if i != len(filenames)-1 {
				return nil, fmt.Errorf("stdin is only allowed as the last filename")
			}
			b, err = io.ReadAll(os.Stdin)
		} else {
			b, err = os.ReadFile(strings.TrimSpace(fn))
		}
		if err := checkDops(string(b)); err != nil {
			return nil, fmt.Errorf("checkDops err:%v", err)
		}
		m, err := values.MapFromYAML(b)
		if err != nil {
			return nil, fmt.Errorf("yaml Unmarshal err:%v", err)
		}
		if m["spec"] == nil {
			delete(m, "spec")
		}
		ConfigBase.MergeFrom(m)
	}
	return nil, nil
}
func checkDops(s string) error {
	mfs, err := manifest.ParseMultiple(s)
	if err != nil {
		return fmt.Errorf("unable to parse file: %v", err)
	}
	if len(mfs) > 1 {
		return fmt.Errorf("")
	}
	return nil
}
func GenerateManifest(files []string, setFlags []string, logger clog.Logger) ([]manifest.ManifestSet, values.Map, error) {
	var chartWarnings util.Errors
	merged, err := MergeInputs(files, setFlags)
	if err != nil {
		return nil, nil, fmt.Errorf("merge inputs: %v", err)
	}
	if err := validateDubboOperator(merged, logger); err != nil {
		return nil, nil, fmt.Errorf("validateDubboOperator err:%v", err)
	}
	allManifests := map[component.Name]manifest.ManifestSet{}
	for _, comp := range component.AllComponents {
		specs, err := comp.Get(merged)
		if err != nil {
			return nil, nil, fmt.Errorf("get component %v: %v", comp.UserFacingName, err)
		}
		for _, spec := range specs {
			compVals := applyComponentValuesToHelmValues(comp, spec, merged)
		}
	}
	return nil, nil, nil
}

func validateDubboOperator(dop values.Map, logger clog.Logger) error {
	warnings, errs := validation.ParseAndValidateDubboOperator(dop)
	if err := errs.ToErrors(); err != nil {
		return err
	}
	if logger != nil {
		for _, w := range warnings {
			logger.LogAndErrorf("%s %v", "❗", w)
		}
	}
	return nil
}

func applyComponentValuesToHelmValues(comp component.Component, spec apis.MetadataCompSpec, merged values.Map) values.Map {
	root := comp.HelmTreeRoot
	if spec.Namespace != "" {
		spec.Namespace = "dubbo-system"
	}
	return merged
}
