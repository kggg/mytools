package main

import (
	"fmt"
	"strings"

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
			finfo.Append(cmd.Words, cmd.Save)
		case "unshift":
			finfo.Unshift(cmd.Words, cmd.Save)
		case "shift":
			finfo.Shift(cmd.Words, cmd.Save)
		case "pop":
			finfo.Pop(cmd.Words, cmd.Save)
		case "delete":
			finfo.Delete(cmd.Words, cmd.Save)
		case "search":
			finfo.Search(cmd.Words)
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
			src := strings.Split(cmd.Words, " ")[0]
			dst := strings.Split(cmd.Words, " ")[1]
			err := cliet.Sendfile(src, dst)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("发送文件%s成功\n", src)
		case "getfile":
			src := strings.Split(cmd.Words, " ")[0]
			dst := strings.Split(cmd.Words, " ")[1]
			err := cliet.Getfile(src, dst)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("接收文件%s成功\n", src)
		default:
			result, err := cliet.Run(cmd.Words)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(result))

		}
	case "web":
		web.Service(cmd.Target)
	}
}
