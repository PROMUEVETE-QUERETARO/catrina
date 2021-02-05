package internal

import "fmt"

func setupWizard(r string) (config Config, err error) {
	const exitMsj = "(type 'exit' to close)"

	config.Port = DefaultPort

	if r != "y" {
		return
	}

	fmt.Printf("Set deploy path:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildPath = r

	fmt.Printf("Set final javascript filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildJS = r

	fmt.Printf("Set final css filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.BuildCSS = r

	fmt.Printf("Set path of input javascript filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.MainJS = r

	fmt.Printf("Set path of input css filename:%v\n", exitMsj)
	_, err = fmt.Scan(&r)
	if err != nil || r == "exit" {
		return
	}
	config.MainCSS = r

	fmt.Println("Set port of trial server?:(y/n)")
	_, err = fmt.Scan(&r)
	if err != nil {
		return
	}
	if r == "y" {
		fmt.Print("Port: ")
		_, err = fmt.Scan(&r)
		if err != nil {
			return
		}
		config.Port = fmt.Sprintf(":%v", r)
	}

	return
}
