package internal

import (
	c "github.com/otiai10/copy"
	"os"
	"path"
)

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
