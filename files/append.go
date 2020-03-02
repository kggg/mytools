package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Append append str to end of line.
func (c *Fileinfo) append(str string) {
	file, mode, err := c.getcontent()
	if err != nil {
		fmt.Printf("getcontent read error: %v\n", err)
		return
	}
	defer file.Close()

	var newcontent []byte
	content := bufio.NewReader(file)
	for {
		line, err := content.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read line: ", err)
			return
		}

		line = bytes.TrimSuffix(line, []byte("\r\n"))
		line = bytes.TrimSuffix(line, []byte("\n"))
		if len(line) == 0 {
			continue
		}
		if c.Save {
			newcontent = append(newcontent, line...)
			newcontent = append(newcontent, []byte(" "+str+"\n")...)
		} else {
			fmt.Printf("%s %s\n", string(line), str)
		}
	}
	if c.Save {
		err = c.save(newcontent, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("save success")
	}
}

//Pop 截取行尾str的内容
func (c *Fileinfo) pop(str string) {
	file, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var newcontent []byte
	content := bufio.NewReader(file)
	for {
		line, err := content.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		//line = bytes.TrimSuffix(line, []byte("\r\n"))
		line = bytes.TrimSuffix(line, []byte("\n"))
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		line = bytes.TrimSuffix(line, []byte(str))
		if c.Save {
			newcontent = append(newcontent, line...)
			newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s\n", line)
		}
	}
	if c.Save {
		err = c.save(newcontent, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("save success")
	}
}
