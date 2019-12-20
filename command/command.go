package command

import(
	"flag"
)

var (
	module string
	target string
	words string
	action string

)

type Cmd struct {
	Module string
	Target string
	Words string
	Action string

}

func init(){
	flag.StringVar(&module, "module", "default", "select module")
	flag.StringVar(&target, "target", "","target action")
	flag.StringVar(&words, "word", "", "strings of words")
	flag.StringVar(&action, "action", "","action")
}

func Command() *Cmd {
	flag.Parse()
	var cmd = &Cmd{
		Module: module,
		Target: target,
		Words: words,
		Action: action,
	}
	return cmd
}