package command

import (
	flag "github.com/spf13/pflag"
)

var (
	module string
	target string
	words  string
	action string
	save   bool
)

type Cmd struct {
	Module string
	Target string
	Words  string
	Action string
	Save   bool
}

func init() {
	flag.StringVarP(&module, "module", "m", "ssh", "选择模块")
	flag.StringVarP(&target, "target", "t", "", "目标对象")
	flag.StringVarP(&words, "word", "w", "", "需要操作的内容")
	flag.StringVarP(&action, "action", "a", "", "操作行为")
	flag.BoolVarP(&save, "save", "s", false, "保存或者不保存，默认不保存, true|false")
}

func Command() *Cmd {
	flag.Parse()
	var cmd = &Cmd{
		Module: module,
		Target: target,
		Words:  words,
		Action: action,
		Save:   save,
	}
	return cmd
}
