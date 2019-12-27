package main

import (
	"fmt"

	"mytools/command"
	"mytools/files"
	"mytools/remote"
	"mytools/remote/config"
	"mytools/web"
)

func main() {
	cmd := command.Command()
	switch cmd.Module {
	case "files":
		finfo := &files.Fileinfo{Filename: cmd.Target}
		switch cmd.Action {
		case "append":
			finfo.Append(cmd.Words)
		case "unshift":
			finfo.Unshift(cmd.Words)
		case "shift":
			finfo.Shift(cmd.Words)
		case "pop":
			finfo.Pop(cmd.Words)
		case "delete":
			finfo.Delete(cmd.Words)
		default:
			fmt.Println("nothing to do")
		}
	case "ssh":
		hostinfo, err := config.Readconfig(cmd.Target)
		if err != nil {
			fmt.Println(err)
			return
		}
		cliet, err := remote.NewClient(hostinfo.Ipaddr, hostinfo.User, hostinfo.Pass, hostinfo.Port, hostinfo.Skey)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch cmd.Action {
		case "sendfile":
		case "getfile":
		default:
			result, err := cliet.Run(cmd.Words)
			if err != nil {
				fmt.Println(err)
				//return
			}
			fmt.Println(string(result))

		}
	case "web":
		web.Service(cmd.Target)
	}
}
