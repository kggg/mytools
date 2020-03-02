package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
)

// Search 查找内容
func (c *Fileinfo) search(str string) {
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
		if bytes.Contains(line, []byte(str)) {
			fmt.Printf("%s", line)
		}
	}

}

// Rsearch 按正则规则pattern查找
func (c *Fileinfo) rsearch(pattern string) {
	reg := regexp.MustCompile(pattern)
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
		ss := reg.FindAll(line, -1)
		if len(ss) > 0 {
			for _, word := range ss {
				fmt.Println(string(word))
			}
		}
	}
}
