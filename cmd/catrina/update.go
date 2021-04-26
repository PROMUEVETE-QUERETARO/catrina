package main

import (
	"fmt"
	"github.com/PROMUEVETE-QUERETARO/catrina/internal"
	"os/exec"
)

func main() {
	v, err := exec.Command("catrina", "version").Output()
	if err != nil {
		fmt.Println("Fatal Error!:", err)
		return
	}
	version := string(v)[:len(string(v))-1]
	err = internal.UpdateCatrina(version, internal.UrlUpdate)
	if err != nil {
		fmt.Println("Fatal Error!:", err)
	}
}
