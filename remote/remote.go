package remote

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

// Client SSH结构体
type Client struct {
	SshClient *ssh.Client
	Hostname  string
	Addr      string
}

// NewClient 创建ssh连接
func NewClient(ip, user, pass string, port int, hostname string, skey bool) (*Client, error) {
	var authMethod ssh.AuthMethod
	authMethod = ssh.Password(pass)
	if skey {
		file := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		auth, err := PublicKeyFile(file)
		if err != nil {
			return nil, err
		}
		authMethod = auth
	}

	sshclient, err := connect(ip, user, pass, port, authMethod)
	if err != nil {
		return nil, err
	}
	var conclient = &Client{}
	conclient.SshClient = sshclient
	conclient.Hostname = hostname
	conclient.Addr = ip
	return conclient, nil
}

func (this *Client) session() *ssh.Session {
	session, err := this.SshClient.NewSession()
	if err != nil {
		return nil
	}
	return session
}

func (this *Client) Run(cmd string) ([]byte, error) {
	session := this.session()
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
	}

	// connet to ssh
	addr := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	sshConn, chans, reqs, err := ssh.NewClientConn(conn, ip, clientConfig)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(sshConn, chans, reqs)

	return client, nil
}
