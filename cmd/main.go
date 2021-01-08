package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jonhteper/go-addons/addons"
	"github.com/jonhteper/go-addons/core"
	c "github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	END_EXPORT        = "//@stop"
	START             = "new"
	RUN_SERVER        = "run"
	BUILD             = "build"
	CONFIG_FILE       = "catrina.config.json"
	DEFAULT_PORT      = ":9095"
	COMPILE_FILE_JS   = "catrina.js"
	EXPORTS_FILE_PATH = "./lib/exports.js"
	FONTS_RELATION    = "./lib/css-fonts-relation.json"
)

type Config struct {
	Port      string `json:"serverPort"`   // port of proof server.
	MainJS    string `json:"inputFileJS"`  // input file javascript location.
	MainCSS   string `json:"inputFileCSS"` // input file css location.
	BuildPath string `json:"deployPath"`   // path where final files will build and where start the proof server.
	BuildJS   string `json:"finalFileJS"`  // final javascript filename.
	BuildCSS  string `json:"finalFileCSS"` // final css filename.
}

func (config *Config) Set(file *os.File) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}

type Import struct {
	names []string
	path  string
}

type ImportCSS struct {
	path  string
	lines []int
}

type RelationCSSFont struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func compileCatrinaJS(exports *os.File) ([]Import, error) {
	scanner := bufio.NewScanner(exports)
	var filesPaths []string
	var directory []Import

	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "export") {
			s := strings.Split(line, " ")
			imp := Import{
				names: strings.Split(s[1][1:len(s[1])-1], ","),
				path:  path.Clean(s[3][1 : len(s[3])-1]),
			}
			filesPaths = append(filesPaths, imp.path)
			directory = append(directory, imp)
		}
	}

	file, err := os.OpenFile(path.Join("temp", COMPILE_FILE_JS), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for _, v := range filesPaths {
		data, err := ioutil.ReadFile(path.Join("lib", v))
		if err != nil {
			return nil, err
		}
		_, err = file.Write(data)
		if err != nil {
			_ = file.Close()
			return nil, err
		}

	}

	return directory, nil
}

func getImportsJS(file *os.File, list []string) ([]string, []int) {
	var lines []int
	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "import") {
			s := strings.Split(line, " ")
			names := strings.Split(s[1][1:len(s[1])-1], ",")
			for _, v := range names {
				list = core.SafeAppend(list, v)
			}
			lines = append(lines, n)
		}
		n++
	}

	return list, lines
}

func getImportsCSS(file *os.File, list []string) ([]string, []int) {
	var lines []int
	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "@import") {
			s := strings.Split(line, " ")
			_imp := strings.ReplaceAll(s[1][1:len(s[1])-2], "lib", "")
			imp := path.Join("/", _imp)
			list = core.SafeAppend(list, path.Join("./lib", path.Clean(imp)))
			lines = append(lines, n)
		}
		n++
	}

	return list, lines
}

func evaluateLine(s string, names []string) bool {
	if !strings.Contains(s, "export ") {
		return false
	}

	for _, v := range names {
		if ev := strings.Contains(s, v); ev {
			return true
		}
	}

	return false
}

func writeImportsJS(ref, build *os.File, names []string) error {
	scanner := bufio.NewScanner(ref)
	ev := false
	for scanner.Scan() {
		if line := scanner.Text(); evaluateLine(line, names) || ev {
			ev = true
			if strings.Contains(line, END_EXPORT) {
				ev = false
				continue
			}

			_, err := build.WriteString(fmt.Sprintf("%v\n", line))
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func writeImportsCSSAndCopyFonts(build *os.File, list []string, config Config) error {
	var err error
	var directory []RelationCSSFont
	_, err = addons.ReadJSONFile(FONTS_RELATION, &directory)
	if err != nil {
		return err
	}

	for _, v := range list {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}

		lines := strings.Split(string(data), "\n")
		for i, v := range lines {
			if strings.Contains(v, "@import") {
				lines[i] = ""
			}
		}
		content := strings.Join(lines, "\n")

		_, err = build.WriteString(content)
		if err != nil {
			return err
		}

		for _, font := range directory {
			if font.Name == path.Base(v) {
				err = c.Copy(font.Path, path.Join(config.BuildPath, path.Base(font.Path)))
				if err != nil {
					return err
				}
			}
		}

	}

	return err
}

func writeFinalFileJS(file *os.File, inputFile string, lines []int) error {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	contentLines := strings.Split(string(data), "\n")
	for _, v := range lines {
		contentLines[v-1] = ""
	}
	mainContent := strings.Join(contentLines, "\n")

	_, err = file.WriteString(mainContent)
	if err != nil {
		return err
	}

	return nil
}

func build(config Config) error {
	var err error
	_ = os.Mkdir("temp", 0755)
	exportsFile, err := os.Open(EXPORTS_FILE_PATH)
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

	imports, lines := getImportsJS(inputFile, []string{})
	var files []string
	for _, v := range imports {
		for _, imp := range directory {
			for _, name := range imp.names {
				if name == v {
					files = core.SafeAppend(files, imp.path)
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

	catrinaJS, err := os.Open(path.Join("temp", COMPILE_FILE_JS))
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
		imports = core.SafeAppend(imports, v)
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

func setupWizard(r string) (config Config, err error) {
	const exitMsj = "(type 'exit' to close)"

	config.Port = DEFAULT_PORT

	if r != "y" {
		return
	}

	fmt.Printf("Set deploy path:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildPath = r

	fmt.Printf("Set final javascript filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildJS = r

	fmt.Printf("Set final css filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildCSS = r

	fmt.Printf("Set path of input javascript filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.MainJS = r

	fmt.Printf("Set path of input css filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.MainCSS = r

	fmt.Println("Set port of trial server?:(y/n)")
	_, err = fmt.Scan(&r)
	if err != nil {
		return
	}
	if r == "y" {
		fmt.Print("Port: ")
		_, err = fmt.Scan(&r)
		if err != nil {
			return
		}
		config.Port = fmt.Sprintf(":%v", r)
	}

	return
}

func newProject(name string) (projectPath string, config Config, err error) {
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

func runServer() error {
	return nil
}

func readConfig() (config Config, err error) {
	_, err = addons.ReadJSONFile(CONFIG_FILE, &config)
	return
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("Not enough args, use '%v', '%v' or '%v' \n", START, RUN_SERVER, BUILD)
		return
	}

	order := args[0]
	if order == START {
		if len(args) < 2 {
			fmt.Println("write a name after 'new'. Example: 'catrina new myProject'")
			return
		}

		projectPath, config, err := newProject(args[1])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		file, err := os.OpenFile(path.Join(projectPath, CONFIG_FILE), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			return
		}
		defer file.Close()

		err = config.Set(file)
		if err != nil {
			return
		}

		fmt.Printf("\nYour configuration is.\n "+
			"Deploy path: %v\n "+
			"Final javascript filename: %v\n "+
			"Final css filename: %v\n "+
			"Input javascript file: %v\n "+
			"Input css file: %v\n "+
			"Server port: %v\n",
			config.BuildPath,
			config.BuildJS,
			config.BuildCSS,
			config.MainJS,
			config.MainCSS,
			config.Port,
		)
		fmt.Printf("\nYou can edit this configuration in file %v\n", CONFIG_FILE)

	} else if order == RUN_SERVER {
		//TODO Crear un servidor en el puerto elegido en la configuración, solo sirve el index.html
		config, err := readConfig()
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}
		fmt.Println(config)
	} else if order == BUILD {
		config, err := readConfig()
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}

		err = build(config)
		if err != nil {
			_ = os.RemoveAll("temp")
			fmt.Println("Fatal Error!:", err)
			return
		}
		fmt.Println("Built!")
	} else {
		fmt.Printf("Not correct args, use '%v', '%v' or '%v' \n", START, RUN_SERVER, BUILD)
		return
	}
	//TODO commands for:
	// -- install js libraries
	// -- update catrina standard lib
	// -- update catrina entire
}
