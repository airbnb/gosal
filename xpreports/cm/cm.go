package cm

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Facts stores output of management tool 'facts'
type Facts map[string]interface{}

// GetFacts will exec into the Facts interface
func GetFacts(confTool, cmdPath, cmdArgs string) (Facts, error) {
	args := strings.Split(cmdArgs, " ")
	cmd := exec.Command(cmdPath, args...)

	o, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "exec facts")
	}

	var facts Facts

	if err := json.Unmarshal(o, &facts); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling Facts")
	}

	switch confTool {
	case "puppet":
		facts = facts["values"].(map[string]interface{})
	case "salt":
		facts = facts["local"].(map[string]interface{})
	}

	facts = FlattenMapStringInterface(facts)

	return facts, nil
}

// FlattenMapStringInterface takes a Facts and returns a flattened Facts
// with no null values
func FlattenMapStringInterface(m Facts) Facts {
	o := make(Facts)
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := FlattenMapStringInterface(child)
			for nk, nv := range nm {
				o[k+"=>"+nk] = nv
			}
		default:
			// Delete null keys
			if v == nil {
				delete(m, k)
				break
			}
			o[k] = v
		}
	}
	return o
}
