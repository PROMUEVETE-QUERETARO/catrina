package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func setupWizard(project string) (err error) {
	var outCommand []byte
	outCommand, err = exec.Command(fmt.Sprintf("%v", path.Join(ExecutablePath(), "./tools/wizard")), project).Output()
	if err != nil {
		return
	}

	out := string(outCommand)
	if strings.Contains(out, "error") || strings.Contains(out, "Error") {
		return errors.New("error executing wizard")
	}

	fmt.Println(out)

	return
}
