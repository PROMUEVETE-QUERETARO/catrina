package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jonhteper/go-addons/addons"
	c "github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	END_EXPORT   = "//@stop"
	START        = "new"
	RUN_SERVER   = "run"
	BUILD        = "build"
	CONFIG_FILE  = "catrina.config.json"
	DEFAULT_PORT = ":9095"
)

type Config struct {
	Port      string `json:"port"`      // Puerto del servidor de pruebas.
	MainJS    string `json:"mainJS"`    // Ruta del archivo principal javascript.
	BuildPath string `json:"buildPath"` // Ruta donde se construirá el archivo final javascript.
	BuildName string `json:"buildName"` // Nombre del archivo final.
}

type Import struct {
	names []string
	path  string
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

func searchFunction(reference, build *os.File, names []string) error {
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

func getImports(file, buildFile *os.File, buildFilePath string) (lines []int, err error) {
	scanner := bufio.NewScanner(file)
	var imports []Import
	n := 1

	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "import") {
			if match, err := regexp.MatchString(`^import(.*);$`, line); err != nil || !match {
				_ = os.Remove(buildFilePath)
				return lines, err
			}

			s := strings.Split(line, " ")
			imp := Import{
				names: strings.Split(s[1][1:len(s[1])-1], ","),
				path:  s[3][2 : len(s[3])-2],
			}

			imports = append(imports, imp)
			lines = append(lines, n)
		}
		n++
	}

	for _, v := range imports {
		f, err := os.Open(v.path)
		if err != nil {
			return lines, err
		}

		_ = searchFunction(f, buildFile, v.names)
		f.Close()
	}

	return
}

func writeFinalFileJS(file *os.File, mainPath string, lines []int) error {
	data, err := ioutil.ReadFile(mainPath)
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
	finalJS := fmt.Sprintf("%v/%v", config.BuildPath, config.BuildName)
	_ = os.Remove(finalJS)

	file, err := os.Open(config.MainJS)
	if err != nil {
		return err
	}
	defer file.Close()

	_ = os.Mkdir("temp", 0755)
	buildFilePath := "./temp/main.build.js"
	_ = os.Remove(buildFilePath)
	buildFile, err := os.OpenFile(buildFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer buildFile.Close()

	lines, err := getImports(file, buildFile, buildFilePath)
	if err != nil {
		return err
	}

	err = writeFinalFileJS(buildFile, buildFilePath, lines)
	if err != nil {
		return err
	}

	err = os.Rename(buildFilePath, finalJS)
	if err != nil {
		return err
	}

	_ = os.Remove("temp")
	return nil
}

func setupWizard(r string) error {
	file, err := os.OpenFile(CONFIG_FILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	config := Config{Port: DEFAULT_PORT}

	if r == "n" {
		// TODO setConfig(config)
	}

	fmt.Println("Set Deploy Path: (type 'exit' to close) ")
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		// TODO setConfig(config)
		return err
	}
	config.BuildPath = r
	fmt.Println(config)

	// TODO setConfig(config)
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

	return setupWizard(r)
}

func runServer() error {
	return nil
}

func readConfig() (config Config, err error) {
	_, err = addons.ReadJSONFile(CONFIG_FILE, &config)
	if err != nil {
		return
	}

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
		//TODO Copiar las librerías y crear el archivo de configuración
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
		fmt.Println(config)
		// TODO verificar que se pueda leer la configuración
		/*err = build(config)
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}*/
	} else {
		fmt.Printf("Not correct args, use '%v', '%v' or '%v' \n", START, RUN_SERVER, BUILD)
		return
	}

}
