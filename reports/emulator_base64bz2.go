package reports

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dsnet/compress/bzip2"
	"github.com/groob/plist"
	"github.com/pkg/errors"
)

type basereport struct {
	AvailableDiskSpace int
	ConsoleUser        string
	OSFamily           string
	MachineInfo        *MachineInfo
	Facter             PuppetFacts
}

// BuildBase64bz2Report will return a compressed and encoded string of our report struct
func BuildBase64bz2Report() (string, error) {

	// gets the absolute path to the config file
	// TODO clean this up, loadconfig should do this work
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", errors.Wrap(err, "bz2: could not determine absolute path to config file")
	}

	s := filepath.Join(dir, "config.json")
	conf, err := LoadConfig(s)
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to load config file")
	}

	// switch for whats "facts" we would send sal
	var facts map[string]interface{}
	switch conf.Management.Tool {
	case "puppet":
		facts, err = GetPuppetFacts()
		if err != nil {
			return "", errors.Wrap(err, "puppet switch: failed to get facts")
		}
	case "chef":
		fmt.Println("although perfectly normal, we dont support chef yet")
	case "salt":
		fmt.Println("people who run salt on the client are strange")
	}

	cDrive, err := GetCDrive()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting c: drive")
	}

	machineInfo, err := EmulateMachineInfo()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting machine info")
	}

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting computer system")
	}

	// TODO report struct needs to switch based on config management tool
	// Facter would change to w/e Sal supported as an input
	report := basereport{
		AvailableDiskSpace: cDrive.FreeSpace,
		MachineInfo:        machineInfo,
		ConsoleUser:        strings.Split(computerSystem.UserName, "\\")[1],
		OSFamily:           "Windows",
		Facter:             facts,
	}

	encodedReport, err := report.CompressAndEncode()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to compress and encode report")
	}

	return encodedReport, nil
}

func (r *basereport) CompressAndEncode() (string, error) {

	var buf bytes.Buffer

	bzw, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestSpeed})
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to bzip2")
	}
	defer bzw.Close()

	enc := plist.NewEncoder(bzw)
	enc.Indent("  ")

	if err = enc.Encode(r); err != nil {
		return "", errors.Wrap(err, "bz2: failed to encode plist")
	}
	bzw.Close()

	report := base64.StdEncoding.EncodeToString(buf.Bytes())
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to base64 encode the report string")
	}

	return report, nil
}

func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "bz2: failed to load config file")
	}

	var conf Config
	if err = json.Unmarshal(file, &conf); err != nil {
		return nil, errors.Wrap(err, "bz2: failed to unmarshal config")
	}

	return &conf, nil
}

// Config will extract the configuration management tool to use.
type Config struct {
	Management *Management
}

// Management is the nested config
type Management struct {
	Tool    string
	Path    string
	Command string
}
