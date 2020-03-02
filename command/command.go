package command

import (
	"errors"
	"mytools/files"
	"mytools/remote"
	"mytools/remote/config"
	"mytools/web"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	path    string
	host    string
	port    int
	cmd     string
	word    string
	action  string
	save    bool
	fileSet *flag.FlagSet
	sshSet  *flag.FlagSet
	webSet  *flag.FlagSet
)

func init() {
	fileSet = flag.NewFlagSet("file", flag.ExitOnError)
	fileSet.StringVarP(&path, "path", "p", "", "文件路径, filepath")
	fileSet.StringVarP(&action, "action", "a", "", "操作行为, operation")
	fileSet.StringVarP(&word, "word", "w", "", "需要操作的内容word")
	fileSet.BoolVarP(&save, "save", "s", false, "操作文件时， 是否保存操作结果，默认不保存, true|false")

	sshSet = flag.NewFlagSet("ssh", flag.ExitOnError)
	sshSet.StringVarP(&host, "host", "h", "localhost", "远程主机名， 在配置文件中设置")
	sshSet.StringVarP(&action, "action", "a", "", "远程执行的模块")

	webSet = flag.NewFlagSet("web", flag.ExitOnError)
	webSet.StringVarP(&host, "host", "h", "127.0.0.1:8080", "WEB服务地址")

}

func Command() error {
	if len(os.Args) < 2 {
		return errors.New("参数太少")
	}
	switch os.Args[1] {
	case "file":
		err := fileSet.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if fileSet.NFlag() < 2 {
			return errors.New("参数太少， 需要指定文件名和操作内容")
		}
		fileinfo := &files.Fileinfo{Filename: path, Save: save, Args: fileSet.Args()}
		err = fileinfo.Run(action, word)
		if err != nil {
			return err
		}
	case "ssh":
		err := sshSet.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if sshSet.NFlag() < 2 {
			return errors.New("参数太少， 需要指定远程主机地址和执行命令")
		}
		hostinfo, err := config.Readconfig(host)
		if err != nil {
			return err
		}
		client := &remote.Client{Addr: hostinfo.Ipaddr, User: hostinfo.User, Pass: hostinfo.Pass,
			Port: hostinfo.Port, Skey: hostinfo.Skey, Args: sshSet.Args()}
		if err := client.RemoteCmd(action); err != nil {
			return err
		}
	case "web":
		err := webSet.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		web.Service(host)
	default:
		return errors.New("可选三个模块: file | ssh | web")
	}
	return nil
}
