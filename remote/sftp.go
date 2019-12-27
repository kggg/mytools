package remote

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/pkg/sftp"
)

// Getfile 从远程机子上复制文件到本地
func (c *Client) Getfile() {

}

// Sendfile 发送本地文件到远程机子上
func (c *Client) Sendfile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	remotename := path.Base(src)
	sftpClient, err := sftp.NewClient(c.SshClient)
	if err != nil {
		return err
	}
	defer sftpClient.Close()
	dstFile, err := sftpClient.Create(path.Join(dst, remotename))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if n == 0 {
			break
		}
		_, err = dstFile.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}
