package internal

import (
	"bufio"
	"fmt"
	c "github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// RelationCSSFont represents the relationship between css files and font files
type RelationCSSFont struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// getImportsCSS get all imports from some css file
func getImportsCSS(file *os.File, list []string) ([]string, []int) {
	var lines []int
	scanner := bufio.NewScanner(file)
	n := 1
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "@import") {
			s := strings.Split(line, " ")
			_imp := strings.ReplaceAll(s[1][1:len(s[1])-2], "lib", "")
			imp := path.Join("/", _imp)
			list = safeAppend(list, path.Join("./lib", path.Clean(imp)))
			lines = append(lines, n)
		}
		n++
	}

	return list, lines
}

func writeImportsCSSAndCopyFonts(build *os.File, list []string, config Config) error {
	var err error
	var directory []RelationCSSFont
	_, err = readJSONFile(FontsRelation, &directory)
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

func GetCSSRule(file *os.File, name string) (content []byte) {
	scanner := bufio.NewScanner(file)
	ev := false
	var rules []string
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, fmt.Sprintf("%v {", name)) ||
			strings.Contains(line, fmt.Sprintf("%v,", name)) || strings.Contains(line, fmt.Sprintf("%v{", name)) || ev {
			ev = true

			rules = append(rules, line)

			if strings.Contains(line, "}") {
				break
			}
		}
	}

	if len(rules) == 0 {
		return
	}

	rule := strings.Join(rules, "\n")
	content = append(content, rule...)
	return
}
