package internal

import (
	"fmt"
	c "github.com/otiai10/copy"
	"os"
	"path"
)

// NewProject create a new project with catrina rules.
func NewProject(name string) (err error) {
	if err = os.Mkdir(name, 0755); err != nil {
		if !os.IsExist(err) {
			return
		}

		return fmt.Errorf("the project %v exist, try with a different name", name)
	}

	projectPath := path.Join(StartPath(), name)
	if err = c.Copy(path.Join(ExecutablePath(), "lib"), path.Join(projectPath, "lib")); err != nil {
		return
	}

	fmt.Print("The project has been created successfully!\n\n Do you want to start the setup wizard?(y/n)")
	var r string
	if _, err = fmt.Scan(&r); err != nil {
		return
	}

	if r == "y" {
		if err = setupWizard(name); err != nil {
			return
		}
	}

	return
}
