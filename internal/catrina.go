package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// BuildProject packages the project as defined in the configuration file.
func BuildProject(config Config) error {
	var err error
	_ = os.Mkdir("temp", 0755)
	exportsFile, err := os.Open(ExportsFilePath)
	if err != nil {
		return err
	}
	defer exportsFile.Close()

	directory, err := compileCatrinaJS(exportsFile)
	if err != nil {
		return err
	}

	inputFile, err := os.Open(config.MainJS)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	imports, lines := getImportsJSInputFile(inputFile)
	var files []string
	for _, v := range imports {
		for _, imp := range directory {
			for _, name := range imp.names {
				if name == v {
					files = safeAppend(files, imp.path)
				}
			}
		}
	}

	for _, file := range files {
		f, err := os.Open(path.Join("lib", file))
		if err != nil {
			return err
		}
		imports, _ = getImportsJS(f, imports)
		_ = f.Close()
	}

	finalJS := path.Join(config.BuildPath, config.BuildJS)
	_ = os.Remove(finalJS)

	buildFilePath := "./temp/main.build.js"
	_ = os.Remove(buildFilePath)
	buildFile, err := os.OpenFile(buildFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer buildFile.Close()

	catrinaJS, err := os.Open(path.Join("temp", CompileFileJs))
	if err != nil {
		return err
	}
	defer catrinaJS.Close()

	err = writeCatrinaCoreJS(buildFile)
	if err != nil {
		return err
	}

	err = writeImportsJS(catrinaJS, buildFile, imports)
	if err != nil {
		return err
	}

	err = writeFinalFileJS(buildFile, config.MainJS, lines)
	if err != nil {
		return err
	}

	err = os.Rename(buildFilePath, finalJS)
	if err != nil {
		return err
	}

	inputFileCSS, err := os.Open(config.MainCSS)
	if err != nil {
		return err
	}
	defer inputFileCSS.Close()

	imports, lines = getImportsCSS(inputFileCSS, []string{})

	buildFilePathCSS := "./temp/styles.build.css"

	buildFileCSS, err := os.OpenFile(buildFilePathCSS, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer buildFileCSS.Close()

	var imp []string
	for _, v := range imports {
		f, err := os.Open(v)
		if err != nil {
			return err
		}

		imp, _ = getImportsCSS(f, imp)

		_ = f.Close()
	}

	for _, v := range imp {
		imports = safeAppend(imports, v)
	}

	err = writeImportsCSSAndCopyFonts(buildFileCSS, imports, config)
	if err != nil {
		return err
	}

	err = writeFinalFileJS(buildFileCSS, config.MainCSS, lines)
	if err != nil {
		return err
	}

	err = os.Rename(buildFilePathCSS, path.Join(config.BuildPath, config.BuildCSS))

	_ = os.RemoveAll("temp")
	return err
}

// NewProject create a new project with catrina rules.
func NewProject(name string) (projectPath string, config Config, err error) {
	startDir, err := os.Getwd()
	if err != nil {
		return
	}

	binDir, err := os.Executable()
	if err != nil {
		return
	}

	err = os.Mkdir(name, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return
		}

		return projectPath, config, fmt.Errorf("the project %v exist, try with a different name", name)
	}

	binPath := path.Dir(binDir)
	projectPath = path.Join(startDir, name)

	err = c.Copy(path.Join(binPath, "lib"), path.Join(projectPath, "lib"))
	if err != nil {
		return
	}

	fmt.Print("The project has been created successfully!\n\n Do you want to start the setup wizard?(y/n)")
	var r string
	_, err = fmt.Scan(&r)
	if err != nil {
		return
	}

	config, err = setupWizard(r)

	return
}

// StartServer run a proof server as defined in the configuration file.
func StartServer(config Config) {
	log.Printf("Listen server in http://localhost%v...", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, http.FileServer(http.Dir(config.BuildPath))))
}

// UpdateCatrina check the latest version available for catrina and update the files if is necessary.
func UpdateCatrina(version, url string) error {
	// Check version
	resp, err := http.Get(fmt.Sprintf("%v?version=%v&os=%v", url, version, runtime.GOOS))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dataUpdate UpdateResponse
	err = json.Unmarshal(data, &dataUpdate)
	if err != nil {
		return err
	}

	if dataUpdate.Error {
		return errors.New(dataUpdate.Msj)
	}

	if !dataUpdate.Update {
		return nil
	}

	// Update
	binDir, err := os.Executable()
	if err != nil {
		return err
	}
	dirBin := path.Dir(binDir)
	updateDir := filepath.Join(dirBin, ".update")

	err = os.Mkdir(updateDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	defer os.RemoveAll(updateDir)

	fmt.Printf("Downloading catrina %v...\n", dataUpdate.Version)
	err = downloadFile(filepath.Join(updateDir, "update.zip"), dataUpdate.Url)
	if err != nil {
		return err
	}

	fmt.Println("Extract files...")
	_, err = unzip(filepath.Join(updateDir, "update.zip"), filepath.Join(updateDir, dataUpdate.Version))
	if err != nil {
		return err
	}

	fmt.Println("Check integrity...")
	sum, err := md5Checksum(filepath.Join(updateDir, dataUpdate.Version, "catrina"))
	if err != nil {
		return err
	}

	if dataUpdate.Checksum != sum {
		return errors.New("the binary is corrupt")
	}

	fmt.Printf("installing %v...\n", dataUpdate.Version)

	err = c.Copy(filepath.Join(updateDir, dataUpdate.Version), dirBin, c.Options{
		Skip: func(src string) (bool, error) {
			if src == "catrina-update" {
				return true, nil
			}
			return false, nil
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("catrina %v is installed\n", dataUpdate.Version)
	return nil
}

// UpdateStandardLib update the standard library files in a project. This action delete the additional
// installed libraries.
func UpdateStandardLib() error {
	projectDir, err := os.Getwd()
	if err != nil {
		return err
	}

	binDir, err := os.Executable()
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(projectDir, "lib"))
	if err != nil {
		return err
	}

	err = os.Mkdir("lib", 0755)
	if err != nil {
		return err
	}

	return c.Copy(path.Join(path.Dir(binDir), "lib"), path.Join(projectDir, "lib"))
}
