package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Fileinfo 文件结构
type Fileinfo struct {
	Filename string
}

// Append append str to end of line.
func (c *Fileinfo) Append(str string) {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		if v == "" || len(v) == 0 || v == "\r\n" {
			continue
		}
		v = strings.Replace(v, "\n", "", 1)
		v = strings.Replace(v, "\r\n", "", 1)
		fmt.Printf("%s %s\n", v, str)
	}
}

// Unshift 在行首添加str
func (c *Fileinfo) Unshift(str string) {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		if v == "" || len(v) == 0 || v == "\r\n" {
			continue
		}
		fmt.Printf("%s %s\n", str, v)
	}
}

//Shift 删除行首中包含str的内容
func (c *Fileinfo) Shift(str string) {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		v = strings.TrimPrefix(v, str)
		fmt.Printf("%s\n", v)
	}
}

//Pop 截取行尾str的内容
func (c *Fileinfo) Pop(str string) {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		v = strings.TrimSuffix(v, str)
		fmt.Printf("%s\n", v)
	}
}

//Delete 删除str的内容
func (c *Fileinfo) Delete(str string) {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range content {
		k := strings.Replace(v, str, "", 1)
		fmt.Printf("%s", k)
	}
}

// getcontent 获取文件的内容
func (c *Fileinfo) getcontent() ([]string, error) {
	content, err := os.Open(c.Filename)
	if err != nil {
		return nil, err
	}
	defer content.Close()
	var data []string
	reader := bufio.NewReader(content)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		data = append(data, string(line))
	}
	return data, nil
}

// Print 打印
func (c *Fileinfo) Print() {
	content, err := c.getcontent()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range content {
		fmt.Println(k, v)
	}
}

// Save 保存操作到文件
func (c *Fileinfo) Save() {

}

// Search 查找内容
func (c *Fileinfo) Search(str string) {

}
