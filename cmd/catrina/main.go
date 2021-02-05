package main

import (
	"catrina/internal"
	"flag"
	"fmt"
	"os"
	"path"
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

		projectPath, config, err := internal.NewProject(args[1])
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		file, err := os.OpenFile(path.Join(projectPath, internal.ConfigFile), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			return
		}
		defer file.Close()

		err = config.Set(file)
		if err != nil {
			return
		}

		fmt.Printf("\nYour configuration is.\n "+
			"Deploy path: %v\n "+
			"Final javascript filename: %v\n "+
			"Final css filename: %v\n "+
			"Input javascript file: %v\n "+
			"Input css file: %v\n "+
			"Server port: %v\n"+
			"\nYou can edit this configuration in file %v\n",
			config.BuildPath,
			config.BuildJS,
			config.BuildCSS,
			config.MainJS,
			config.MainCSS,
			config.Port,
			internal.ConfigFile,
		)
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
			err := internal.UpdateCatrina()
			if err != nil {
				fmt.Println("Fatal Error!:", err)
				return
			}
		} else {
			fmt.Printf("Write 'lib' to update standar library. This action, replace all content of " +
				"directory ./lib .\nWrite 'catrina' to update tool files.\n")
		}
	case internal.UpdateTool:
		err := internal.UpdateCatrina()
		if err != nil {
			fmt.Println("Fatal Error!:", err)
			return
		}
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
	default:
		fmt.Printf("Not correct args, use '%v', '%v' or '%v' \n",
			internal.Start,
			internal.RunServer,
			internal.Build,
		)
	}
}
