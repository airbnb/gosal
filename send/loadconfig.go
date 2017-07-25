package send

import (
  	"io/ioutil"
    "log"
    "encoding/json"
)

//loader for sal configuration takes path to config as argument

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}

type Configuration struct {
	Key string
  URL string
}
