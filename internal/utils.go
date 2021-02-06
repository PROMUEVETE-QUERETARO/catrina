package internal

import (
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version         = "v1.1.0-beta.2"
	UrlUpdate       = "https://apps.promuevete.mx/catrina/version.php"
	EndExport       = "//@stop"
	Start           = "new"
	RunServer       = "run"
	Build           = "build"
	Update          = "update"
	UpdateTool      = "upgrade"
	VersionTool     = "version"
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

func downloadFile(filepath, url string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func unzip(filePath, outputPath string) (filenames []string, err error) {
	// function copied from https://golangcode.com/unzip-files-in-go/
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return
	}
	defer r.Close()

	for _, file := range r.File {
		fpath := filepath.Join(outputPath, file.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(outputPath)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if file.FileInfo().IsDir() {
			// Make Folder
			_ = os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := file.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	return
}

func md5Checksum(path string) (sum string, err error) {
	// function copied from https://ispycode.com/Blog/golang/2016-10/Md5-checksum
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil

}

type UpdateResponse struct {
	Update   bool   `json:"update"`
	Version  string `json:"version"`
	Url      string `json:"url"`
	Checksum string `json:"checksum"`
	Error    bool   `json:"error"`
	Msj      string `json:"msj"`
}
