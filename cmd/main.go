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
)

type Config struct {
	Port      string `json:"serverPort"` // Puerto del servidor de pruebas.
	MainJS    string `json:"finalFile"`  // Ruta del archivo principal javascript.
	BuildPath string `json:"deployPath"` // Ruta donde se construirá el archivo final javascript.
	BuildName string `json:"inputFile"`  // Nombre del archivo final.
}

func (config *Config) Set(file *os.File) error {
	data, err := json.Marshal(config)
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

func getImports(file *os.File, list []string) ([]string, []int) {
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

func writeImports(reference, build *os.File, names []string) error {
	scanner := bufio.NewScanner(reference)
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

	imports, lines := getImports(inputFile, []string{})
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
		imports, _ = getImports(f, imports)
		_ = f.Close()
	}

	finalJS := path.Join(config.BuildPath, config.BuildName)
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

	err = writeImports(catrinaJS, buildFile, imports)
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

	_ = os.RemoveAll("temp")
	return err
}

func setupWizard(r, projectPath string) error {
	const exitMsj = "(type 'exit' to close)"
	file, err := os.OpenFile(path.Join(projectPath, CONFIG_FILE), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	config := Config{Port: DEFAULT_PORT}

	if r != "y" {
		err = config.Set(file)
		return err
	}

	fmt.Printf("Set deploy path:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		_ = config.Set(file)
		return err
	}
	config.BuildPath = r

	fmt.Printf("Set final filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		_ = config.Set(file)
		return err
	}
	config.BuildName = r

	fmt.Printf("Set path of input filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		_ = config.Set(file)
		return err
	}
	config.MainJS = r

	fmt.Println("Set port of trial server?:(y/n)")
	_, err = fmt.Scan(&r)
	if err != nil {
		_ = config.Set(file)
		return err
	}
	if r == "y" {
		fmt.Print("Port: ")
		_, err = fmt.Scan(&r)
		if err != nil {
			_ = config.Set(file)
			return err
		}
		config.Port = fmt.Sprintf(":%v", r)
	}

	err = config.Set(file)
	if err != nil {
		return err
	}

	fmt.Printf("\nYour configuration is.\n "+
		"Deploy path: %v\n "+
		"Final filename: %v\n "+
		"Input file: %v\n "+
		"Server port: %v\n",
		config.BuildPath,
		config.BuildName,
		config.MainJS,
		config.Port,
	)
	fmt.Printf("\n You can edit this configuration in file %v\n", CONFIG_FILE)

	return nil
}

func newProject(name string) error {
	startDir, err := os.Getwd()
	if err != nil {
		return err
	}

	binDir, err := os.Executable()
	if err != nil {
		return err
	}

	err = os.Mkdir(name, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
		return fmt.Errorf("the project %v exist, try with a different name", name)
	}

	binPath := path.Dir(binDir)
	projectPath := path.Join(startDir, name)

	err = c.Copy(path.Join(binPath, "lib"), path.Join(projectPath, "lib"))
	if err != nil {
		return err
	}

	fmt.Println("The project has been created successfully!")
	fmt.Println("Do you want to start the setup wizard?(y/n)")
	var r string
	_, err = fmt.Scan(&r)
	if err != nil {
		return err
	}

	return setupWizard(r, projectPath)
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
		err := newProject(args[1])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}
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

}
