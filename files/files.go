package files

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Fileinfo 文件结构
type Fileinfo struct {
	Filename string
	Save     bool
	Args     []string
}

func (c *Fileinfo) Run(action, str string) error {
	switch action {
	case "append":
		c.append(str)
	case "pop":
		c.pop(str)
	case "shift":
		c.shift(str)
	case "unshift":
		c.unshift(str)
	case "delete":
		c.delete(str)
	case "search":
		c.search(str)
	case "rsearch":
		c.rsearch(str)
	case "replace":
		c.replace()
	case "print":
		c.print()
	default:
		return errors.New("Doesn't have operation function of [" + action + "]")
	}
	return nil

}

// Save 保存content到文件c.Filename, mode为文件模式
func (c *Fileinfo) save(content []byte, mode os.FileMode) error {
	err := ioutil.WriteFile(c.Filename, content, mode)
	if err != nil {
		return fmt.Errorf("write file erro: %w\n", err)
	}
	return nil
}

// Unshift 在行首添加str
func (c *Fileinfo) unshift(str string) {
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
		if len(line) == 0 {
			continue
		}
		if c.Save {
			newcontent = append(newcontent, []byte(str)...)
			newcontent = append(newcontent, line...)
			//newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s %s", str, line)
		}
	}
	if c.Save {
		err = c.save(newcontent, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("save unshift opertation success")
	}
}

//Shift 删除行首中包含str的内容
func (c *Fileinfo) shift(str string) {
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
		line = bytes.TrimPrefix(line, []byte(str))
		line = bytes.TrimPrefix(line, []byte(" "))
		if c.Save {
			newcontent = append(newcontent, line...)
			//newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s", line)
		}
	}
	if c.Save {
		err = c.save(newcontent, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("shift success")
	}
}

//Delete 删除str的内容
func (c *Fileinfo) delete(str string) {
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
		line = bytes.Replace(line, []byte(str), []byte(""), 1)
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
		fmt.Println("delete success")
	}
}

// getcontent 获取文件的内容
func (c *Fileinfo) getcontent() (*os.File, os.FileMode, error) {
	f, err := os.Open(c.Filename)
	if err != nil {
		return nil, 0, fmt.Errorf("Can not open file: %w", err)
	}
	//defer f.Close()
	finfo, err := f.Stat()
	if err != nil {
		return nil, 0, fmt.Errorf("read file state error: %w", err)
	}
	mode := finfo.Mode()
	return f, mode, nil
}

// Print 打印
func (c *Fileinfo) print() {
	f, _, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	content := bufio.NewReader(f)
	for {
		line, err := content.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("%s", line)
	}
}

// Replace 替换内容
func (c *Fileinfo) replace() {
	f, _, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	content := bufio.NewReader(f)
	for {
		line, err := content.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		line = bytes.Replace(line, []byte(c.Args[0]), []byte(c.Args[1]), -1)
		fmt.Printf("%s", line)
	}
}
