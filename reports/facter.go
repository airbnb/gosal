package reports

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// PuppetFacts stores output of puppet facts command
type PuppetFacts map[string]interface{}

// GetPuppetFacts will exec into the PuppetFacts interface
func GetPuppetFacts(cmdPath, cmdArgs string) (PuppetFacts, error) {
	// TODO: args should be a parsed slice?
	cmd := exec.Command(cmdPath, cmdArgs)

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return PuppetFacts{}, errors.Wrap(err, "exec puppet facts")
	}

	var pf PuppetFacts

	if err := json.Unmarshal(o, &pf); err != nil {
		return PuppetFacts{}, errors.Wrap(err, "failed unmarshalling Puppet Facts")
	}

	v := pf["values"].(map[string]interface{})

	pf = FlattenMapStringInterface(v)

	return pf, nil
}

// FlattenMapStringInterface takes a map[string]interface and returns a flattened map[string]interface{}
func FlattenMapStringInterface(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := FlattenMapStringInterface(child)
			for nk, nv := range nm {
				o[k+"=>"+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}
