package main

import (
	"fmt"
	"mytools/command"
)

func main() {
	err := command.Command()
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		switch cmd.Module {
		case "files":
		case "ssh":
		case "web":
			web.Service(cmd.Target)
		default:
			fmt.Printf("格式不对")
		}
	*/
}
