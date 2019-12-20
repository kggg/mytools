package files
import(
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
)

type Fileinfo struct {
	Filename string
}

func (c *Fileinfo) Append(str string){
    content, err := c.getcontent()
    if err != nil {
        fmt.Println(err)
        return
	}
    for _,v := range content {
        if v == "" || len(v) == 0 || v =="\r\n" {
            continue
        }
        v = strings.Replace(v, "\n", "", 1)
        v = strings.Replace(v, "\r\n", "", 1)
        fmt.Printf("%s %s\n",v, str)
    }
}

func (c *Fileinfo) Unshift(str string){
    content, err := c.getcontent()
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, v := range content {
        if v == "" || len(v) == 0 || v =="\r\n" {
            continue
        }
        fmt.Printf("%s %s\n",str, v)
    }
}

func (c *Fileinfo) Shift(str string){
    content, err := c.getcontent()
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, v := range content {
        v = strings.TrimPrefix(v, str)
        fmt.Printf("%s\n",  v)
    }
}

func (c *Fileinfo) Pop(str string){
    content, err := c.getcontent()
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, v := range content {
        v = strings.TrimSuffix(v, str)
        fmt.Printf("%s", v)
    }
}

func (c *Fileinfo) Delete(str string){
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

func (c *Fileinfo) getcontent()([]string, error){
    content, err := os.Open(c.Filename)
    if err != nil {
        return nil, err
	}
	defer content.Close()
    var data []string
    reader := bufio.NewReader(content)
    for {
		line,_, err := reader.ReadLine()
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

func (c *Fileinfo) Print(){
    content, err := c.getcontent()
    if err != nil {
        fmt.Println(err)
        return
    }
    for k, v := range content {
        fmt.Println(k, v)
    }
}