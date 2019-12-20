package command

import(
	"flag"
)

var (
	target string
	filepath string
)
func init(){
	flag.StringVar(&target, "target", "default", "target action")
	flag.StringVar(&filepath, "file", "","the file path")
}