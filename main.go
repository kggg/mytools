package main
import(
	"fmt"

	"mytools/command"
	"mytools/files"
)
func main(){
	cmd := command.Command()
	switch cmd.Module {
	case "files":
		finfo := &files.Fileinfo{ Filename: cmd.Target }
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
		fmt.Println("devp...")
	}
}