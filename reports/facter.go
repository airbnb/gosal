package reports

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

// PuppetFacts stores output of puppet facts command
type PuppetFacts map[string]interface{}

// GetPuppetFacts will exec into the PuppetFacts interface
func GetPuppetFacts() (PuppetFacts, error) {
	// gets the absolute path to the config file
	// TODO clean this up, loadconfig should do this work
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return PuppetFacts{}, errors.Wrap(err, "facter: could not determine absolute path to config file")
	}

	s := filepath.Join(dir, "config.json")
	conf, err := LoadConfig(s)
	if err != nil {
		return PuppetFacts{}, errors.Wrap(err, "facter: failed to load config file")
	}

	cmd := exec.Command(conf.Management.Path, conf.Management.Command)

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
