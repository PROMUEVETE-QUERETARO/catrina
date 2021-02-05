package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Import represents the relation between exported functions javascript
// and the file location where this functions has been written
type Import struct {
	names []string
	path  string
}

// compileCatrinaJS gets all the installed javascript code and transcribes it all to a single file
func compileCatrinaJS(exports *os.File) ([]Import, error) {
	scanner := bufio.NewScanner(exports)
	var filesPaths []string
	var directory []Import

	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "export") && !strings.Contains(line, "core.js") {
			s := strings.Split(line, " ")
			imp := Import{
				names: strings.Split(s[1][1:len(s[1])-1], ","),
				path:  path.Clean(s[3][1 : len(s[3])-1]),
			}
			filesPaths = append(filesPaths, imp.path)
			directory = append(directory, imp)
		}
	}

	file, err := os.OpenFile(path.Join("temp", CompileFileJs), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
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

// getImportsJSInputFile get all imports from javascript input file. Only get the objects defined
// in installed libraries.
func getImportsJSInputFile(file *os.File) (list []string, lines []int) {
	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "import") && strings.Contains(line, "lib") {
			s := strings.Split(line, " ")
			names := strings.Split(s[1][1:len(s[1])-1], ",")
			for _, v := range names {
				list = safeAppend(list, v)
			}
			lines = append(lines, n)
		}
		n++
	}

	return
}

// getImportsJS get all imports from some javascript file
func getImportsJS(file *os.File, list []string) ([]string, []int) {
	var lines []int
	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "import") && !strings.Contains(line, "core.js") {
			s := strings.Split(line, " ")
			names := strings.Split(s[1][1:len(s[1])-1], ",")
			for _, v := range names {
				list = safeAppend(list, v)
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

func writeCatrinaCoreJS(build *os.File) error {
	data, err := ioutil.ReadFile(CatrinaCoreJs)
	if err != nil {
		return err
	}

	_, err = build.Write(data)

	return err
}

func writeImportsJS(ref, build *os.File, names []string) error {
	scanner := bufio.NewScanner(ref)
	ev := false
	for scanner.Scan() {
		if line := scanner.Text(); evaluateLine(line, names) || ev {
			ev = true
			if strings.Contains(line, EndExport) {
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
