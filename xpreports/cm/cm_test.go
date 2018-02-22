package cm

import (
	"encoding/json"
	"testing"
)

func TestFlattenMapStringInterface(t *testing.T) {
	jsonTestWithNulls := `{
		"name": "Jimmy",
		"age": 23,
		"test": null,
		"hello": {
			"world": null,
			"goodbye": "Jimmy"
		},
		"alphabet": {
			"one": "alpha",
			"two": "beta",
			"three": "charlie",
			"so_nested": {
				"four": "delta",
				"five": "echo",
				"six": null
			}
		}
	}`

	var testFacts Facts

	if err := json.Unmarshal([]byte(jsonTestWithNulls), &testFacts); err != nil {
		t.Errorf("Unable to unmarshal 'jsonTestWithNulls'")
	}
	o := FlattenMapStringInterface(testFacts)

	null := nullSearcher(o)
	if null == true {
		t.Errorf("Null value found. Please fix 'FlattenMapStringInterface'")
	}
}

// nullSearcher recursively looks for null values. Returning true if any are found.
func nullSearcher(aMap map[string]interface{}) bool {
	out := false
	for _, val := range aMap {
		if val == nil {
			return true
		}
		switch val.(type) {
		case map[string]interface{}:
			out = nullSearcher(val.(map[string]interface{}))
		}
	}
	return out
}
