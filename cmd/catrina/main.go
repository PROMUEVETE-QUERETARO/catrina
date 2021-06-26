package main

import (
	"flag"
	"fmt"
	"github.com/PROMUEVETE-QUERETARO/catrina/internal"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("Not enough args, use '%v', '%v' or '%v' \n",
			internal.Start,
			internal.RunServer,
			internal.Build,
		)
		return
	}

	switch order := args[0]; order {
	case internal.Start:
		if len(args) < 2 {
			fmt.Printf("write a name after '%v'. Example: 'catrina new myProject'\n", internal.Start)
			return
		}

		if err := internal.NewProject(args[1]); err != nil {
			fmt.Println("Error!", err)
			return
		}

	case internal.Update:
		if len(args) < 2 {
			fmt.Printf("Write 'lib' to update standar library. This action, replace all content of " +
				"directory ./lib .\nWrite 'catrina' to update tool files. \n")
			return
		}

		if args[1] == "lib" {
			err := internal.UpdateStandardLib()
			if err != nil {
				fmt.Println("Fatal Error!:", err)
				return
			}
			fmt.Println("the standard catrina library has been updated")
		} else if args[1] == "catrina" {
			fmt.Println("run catrina-update")
		} else {
			fmt.Printf("Write 'lib' to update standar library. This action, replace all content of " +
				"directory ./lib .\nWrite 'catrina' to update tool files.\n")
		}
	case internal.UpdateTool:
		fmt.Println("run catrina-update")
	case internal.RunServer:
		config, err := internal.ReadConfig()
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}
		internal.StartServer(config)
	case internal.Build:
		config, err := internal.ReadConfig()
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}

		err = internal.BuildProject(config)
		if err != nil {
			_ = os.RemoveAll("temp")
			fmt.Println("Fatal Error!:", err)
			return
		}
		fmt.Println("Built!")
	case internal.VersionTool:
		fmt.Println(internal.Version)
	default:
		fmt.Printf("Not correct args, use '%v', '%v', %v or '%v' \n",
			internal.Start,
			internal.RunServer,
			internal.Build,
			internal.VersionTool,
		)
	}
}
