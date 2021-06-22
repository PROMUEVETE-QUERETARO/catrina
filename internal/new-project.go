package internal

import (
	"fmt"
	c "github.com/otiai10/copy"
	"os"
	"path"
)

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
