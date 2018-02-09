package reports

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Facts stores output of puppet & chef facts command
type Facts map[string]interface{}

// GetPuppetFacts will exec into the Facts interface
func GetPuppetFacts(cmdPath, cmdArgs string) (Facts, error) {
	args := strings.Split(cmdArgs, " ")
	cmd := exec.Command(cmdPath, args...)

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return Facts{}, errors.Wrap(err, "exec puppet facts")
	}

	var pf Facts

	if err := json.Unmarshal(o, &pf); err != nil {
		return Facts{}, errors.Wrap(err, "failed unmarshalling Puppet Facts")
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
