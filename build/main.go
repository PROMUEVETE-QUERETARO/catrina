package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const FINAL_EXPORT = "//@stop"

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
			if strings.Contains(line, FINAL_EXPORT) {
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

func main() {
	mainPath := "./src/main.js"
	file, err := os.Open(mainPath)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}
	defer file.Close()

	buildFilePath := "./src/main.build.js"
	_ = os.Remove(buildFilePath)
	buildFile, err := os.OpenFile(buildFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}
	defer buildFile.Close()

	lines, err := getImports(file, buildFile, buildFilePath)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}

	mainContentBytes, err := ioutil.ReadFile(mainPath)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}
	mainContentLines := strings.Split(string(mainContentBytes), "\n")
	for _, v := range lines {
		mainContentLines[v-1] = ""
	}
	mainContent := strings.Join(mainContentLines, "\n")

	_, err = buildFile.WriteString(mainContent)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}

	// Mover archivo final
	err = os.Rename(buildFilePath, "./main.js")
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}
}
