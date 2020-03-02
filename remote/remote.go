package remote

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

// Client SSH结构体
type Client struct {
	SshClient *ssh.Client
	Addr      string
	User      string
	Pass      string
	Port      int
	Skey      bool
	Args      []string
}

// init 初始化Client
func (c *Client) init() (*Client, error) {
	var authMethod ssh.AuthMethod
	authMethod = ssh.Password(c.Pass)
	if c.Skey {
		file := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		auth, err := PublicKeyFile(file)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		authMethod = auth
	}

	sshclient, err := connect(c.Addr, c.User, c.Pass, c.Port, authMethod)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	c.SshClient = sshclient
	return c, nil
}

// RemoteCmd 通过action进行远程操作
func (c *Client) RemoteCmd(action string) error {
	client, err := c.init()
	if err != nil {
		return err
	}
	switch action {
	case "sendfile":
		err := client.Sendfile(c.Args[0], c.Args[1])
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		fmt.Printf("发送文件%s成功\n", c.Args[0])
	case "getfile":
		err := client.Getfile(c.Args[0], c.Args[1])
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("%w", err)
		}
		fmt.Printf("接收文件%s成功\n", c.Args[0])
	case "cmd":
		result, err := client.Run(c.Args[0])
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		fmt.Println(string(result))
	default:
	}
	return nil
}

func (c *Client) session() *ssh.Session {
	session, err := c.SshClient.NewSession()
	if err != nil {
		return nil
	}
	return session
}

func (c *Client) Run(cmd string) ([]byte, error) {
	session := c.session()
	defer session.Close()
	res, err := session.CombinedOutput(cmd)
	return res, err
}

func PublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func connect(ip, user, password string, port int, authMethod ssh.AuthMethod) (*ssh.Client, error) {
	if ip == "" || user == "" {
		return nil, errors.New("Username or IPaddress empty")
	}
	timeout := 30 * time.Second
	clientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{authMethod},
		//	HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	// connet to ssh
	addr := fmt.Sprintf("%s:%d", ip, port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
