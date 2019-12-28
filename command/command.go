package command

import (
	flag "github.com/spf13/pflag"
)

var (
	module string
	target string
	words  string
	action string
)

type Cmd struct {
	Module string
	Target string
	Words  string
	Action string
}

func init() {
	flag.StringVarP(&module, "module","m", "ssh", "选择模块")
	flag.StringVarP(&target, "target", "t", "", "目标对象")
	flag.StringVarP(&words, "word", "w","", "需要操作的内容")
	flag.StringVarP(&action, "action", "a", "", "操作行为")
}

func Command() *Cmd {
	flag.Parse()
	var cmd = &Cmd{
		Module: module,
		Target: target,
		Words:  words,
		Action: action,
	}
	return cmd
}
