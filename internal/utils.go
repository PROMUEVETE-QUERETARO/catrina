package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	EndExport       = "//@stop"
	Start           = "new"
	RunServer       = "run"
	Build           = "build"
	Update          = "update"
	UpdateTool      = "upgrade"
	ConfigFile      = "catrina.config.json"
	DefaultPort     = ":9095"
	CompileFileJs   = "catrina.js"
	CatrinaCoreJs   = "./lib/core/core.js"
	ExportsFilePath = "./lib/exports.js"
	FontsRelation   = "./lib/css-fonts-relation.json"
)

// Config is the catrina's project configuration
type Config struct {
	MainJS    string `json:"inputFileJS"`  // input file javascript location.
	MainCSS   string `json:"inputFileCSS"` // input file css location.
	BuildPath string `json:"deployPath"`   // path where final files will build and where start the proof server.
	BuildJS   string `json:"finalFileJS"`  // final javascript filename.
	BuildCSS  string `json:"finalFileCSS"` // final css filename.
	Port      string `json:"serverPort"`   // port of proof server.
}

// Write the config file
func (config *Config) Set(file *os.File) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}

func safeAppend(slice []string, addon string) []string {
	var verify bool
	for i := 0; i <= len(slice)-1; i++ {

		if slice[i] == addon {
			verify = true
		}

	}
	if !verify {
		slice = append(slice, addon)
		return slice
	}
	return slice
}

func readJSONFile(name string, v interface{}) (content string, err error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}

	content = string(file)
	err = json.Unmarshal(file, v)
	return
}

// ReadConfig read the configuration file
func ReadConfig() (config Config, err error) {
	_, err = readJSONFile(ConfigFile, &config)
	return
}
