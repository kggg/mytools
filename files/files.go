package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

// Fileinfo 文件结构
type Fileinfo struct {
	Filename string
}

// Append append str to end of line.
func (c *Fileinfo) Append(str string, save bool) {
	content, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	var f *os.File
	if save {
		f, err = os.OpenFile(c.Filename, os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}
	var newcontent []byte
	for _, v := range content {
		if len(v) == 0 {
			continue
		}
		v = bytes.TrimSuffix(v, []byte("\n"))
		v = bytes.TrimSuffix(v, []byte("\r\n"))

		if save {
			newcontent = append(newcontent, v...)
			newcontent = append(newcontent, []byte(str+"\n")...)
		} else {
			fmt.Printf("%s %s\n", string(v), str)
		}
	}
	if save {
		_, err = f.Write(newcontent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("append success")
	}
}

// Unshift 在行首添加str
func (c *Fileinfo) Unshift(str string, save bool) {
	content, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	var f *os.File
	if save {
		f, err = os.OpenFile(c.Filename, os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}
	var newcontent []byte
	for _, v := range content {
		if len(v) == 0 {
			continue
		}
		if save {
			newcontent = append(newcontent, []byte(str)...)
			newcontent = append(newcontent, v...)
			newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s %s\n", str, v)
		}
	}
	if save {
		_, err = f.Write(newcontent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("unshift success")
	}
}

//Shift 删除行首中包含str的内容
func (c *Fileinfo) Shift(str string, save bool) {
	content, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	var f *os.File
	if save {
		f, err = os.OpenFile(c.Filename, os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}
	var newcontent []byte
	for _, v := range content {
		v = bytes.TrimPrefix(v, []byte(str))
		if save {
			newcontent = append(newcontent, v...)
			newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s\n", v)
		}
	}
	if save {
		_, err = f.Write(newcontent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("shift success")
	}
}

//Pop 截取行尾str的内容
func (c *Fileinfo) Pop(str string, save bool) {
	content, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	var f *os.File
	if save {
		f, err = os.OpenFile(c.Filename, os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}
	var newcontent []byte
	for _, v := range content {
		v = bytes.TrimSuffix(v, []byte(str))
		if save {
			newcontent = append(newcontent, v...)
			newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s\n", v)
		}
	}
	if save {
		_, err = f.Write(newcontent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("pop success")
	}
}

//Delete 删除str的内容
func (c *Fileinfo) Delete(str string, save bool) {
	content, mode, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	var f *os.File
	if save {
		f, err = os.OpenFile(c.Filename, os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}
	var newcontent []byte
	for _, v := range content {
		v := bytes.Replace(v, []byte(str), []byte(""), 1)
		if save {
			newcontent = append(newcontent, v...)
			newcontent = append(newcontent, []byte("\n")...)
		} else {
			fmt.Printf("%s\n", v)
		}
	}
	if save {
		_, err = f.Write(newcontent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("shift success")
	}
}

// getcontent 获取文件的内容
func (c *Fileinfo) getcontent() ([][]byte, os.FileMode, error) {
	content, err := os.Open(c.Filename)
	if err != nil {
		return nil, 0, err
	}
	defer content.Close()
	finfo, err := content.Stat()
	if err != nil {
		return nil, 0, err
	}
	mode := finfo.Mode()
	var data [][]byte
	reader := bufio.NewReader(content)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		data = append(data, line)
	}
	return data, mode, nil
}

// Print 打印
func (c *Fileinfo) Print() {
	content, _, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range content {
		fmt.Println(k, v)
	}
}

// Search 查找内容
func (c *Fileinfo) Search(str string) {
	content, _, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		if bytes.Contains(v, []byte(str)) {
			fmt.Printf("%s\n", v)
		}
	}

}

func (c *Fileinfo) Rsearch(pattern string) {
	reg := regexp.MustCompile(pattern)
	content, _, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		ss := reg.FindAll(v, -1)
		if len(ss) > 0 {
			for _, word := range ss {
				fmt.Println(string(word))
			}
		}
	}
}

// Replace 替换内容
func (c *Fileinfo) Replace(src, dst string) {

}
