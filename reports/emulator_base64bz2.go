package reports

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}

	s := filepath.Join(dir, "config.json")
	conf, err := loadConfig(s)
	if err != nil {
		log.Fatal(err)
	}

	// switch for whats "facts" we would send sal
	var facts map[string]interface{}
	switch conf.Management.Tool {
	case "puppet":
		facts, err = GetPuppetFacts()
		fmt.Println("got here")
		if err != nil {
			errors.Wrap(err, "puppet switch: failed to get facts")
		}
	case "chef":
		fmt.Println("although perfectly normal, we dont support chef yet")
	case "salt":
		fmt.Println("people who run salt on the client are strange")
	}

	cDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	machineInfo, err := EmulateMachineInfo()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: problem with emulating machine info: %s", err)
	}

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		// TODO return the error here?
		log.Printf("machine info: computer system information failed: %s", err)
	}

	// TODO report struct need to switch based on config management tool
	// Facter would change to w/e Sal supported as an input
	report := basereport{
		AvailableDiskSpace: cDrive.FreeSpace,
		MachineInfo:        machineInfo,
		ConsoleUser:        strings.Split(computerSystem.UserName, "\\")[1],
		OSFamily:           "Windows",
		Facter:             facts,
	}

	// fmt.Println(report)

	encodedReport, err := report.CompressAndEncode()
	if err != nil {
		// TODO return the error here?
		log.Printf("compress and encode failed: %s", err)
	}

	return encodedReport, nil
}

func (r *basereport) CompressAndEncode() (string, error) {

	var buf bytes.Buffer

	bzw, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestSpeed})
	if err != nil {
		log.Fatal(err)
	}
	defer bzw.Close()

	enc := plist.NewEncoder(bzw)
	enc.Indent("  ")

	if err = enc.Encode(r); err != nil {
		log.Fatal(err)
	}
	bzw.Close()

	report := base64.StdEncoding.EncodeToString(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return report, nil
}

func loadConfig(path string) (*Config, error) {
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
