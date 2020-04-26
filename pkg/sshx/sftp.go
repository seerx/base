package sshx

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SFTP 连接
type SFTP struct {
	Client *sftp.Client
}

// NewSFTP 创建 sftp 连接
func NewSFTP(user, password, host string, port int) (*SFTP, error) {
	auth := []ssh.AuthMethod{}
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	cfg := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: hostKeyCallbk,
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	var client *ssh.Client
	var err error
	if client, err = ssh.Dial("tcp", addr, cfg); err != nil {
		return nil, err
	}
	var sftpClient *sftp.Client
	if sftpClient, err = sftp.NewClient(client); err != nil {
		return nil, err
	}
	return &SFTP{
		Client: sftpClient,
	}, nil
}

// IsSendEnd 发送结束
func IsSendEnd(err error) bool {
	return err == io.EOF
}

// Close 断开连接
func (s *SFTP) Close() error {
	if s.Client != nil {
		err := s.Client.Close()
		s.Client = nil
		return err
	}
	return nil
}

// WriteFile 发送文件
func (s *SFTP) WriteFile(destFilePath string, data io.Reader) error {
	dstFile, err := s.Client.Create(destFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, data)
	// _, err = utils.Stream(from, dstFile)
	return err
}

// PutFile 上传文件
func (s *SFTP) PutFile(dest string, fromPath string) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return s.WriteFile(dest, file)
}
