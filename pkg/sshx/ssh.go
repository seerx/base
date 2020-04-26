package sshx

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/seerx/base"
	"golang.org/x/crypto/ssh"
)

// SSH 连接
type SSH struct {
	client              *ssh.Client
	Session             *ssh.Session
	ch                  chan *pipObject
	in                  io.Writer
	currentCommand      string
	currentCommandOK    string
	currentCommandError string
	commandChan         chan *response
}

type response struct {
	data     string
	err      error
	complete bool
}

// IsShellExit 是否是正常结束
func IsShellExit(err error) bool {
	_, ok := err.(*ssh.ExitMissingError)
	return ok
}

// NewSSH  创建 ssh 连接
func NewSSH(user, password, host string, port int) (*SSH, error) {
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
	var ss *ssh.Session
	if ss, err = client.NewSession(); err != nil {
		return nil, err
	}

	return &SSH{
		client:  client,
		Session: ss,
	}, nil
}

// ExecBackground 后台执行
func (s *SSH) ExecBackground(cmd string) error {
	if s.currentCommand != "" {
		return fmt.Errorf("Command %s is running", s.currentCommand)
	}

	cmd = fmt.Sprintf("nohup %s &\n", cmd)
	_, err := s.in.Write([]byte(cmd))
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}

// Exec 运行命令
func (s *SSH) Exec(cmd string) (string, bool, error) {
	if cmd == s.currentCommand {
		return s.ReadNext()
	}
	if s.currentCommand != "" {
		return "", true, fmt.Errorf("Command %s is running", s.currentCommand)
	}
	s.currentCommand = cmd
	s.currentCommandOK = base.MD5s(cmd, "ok")
	s.currentCommandError = base.MD5s(cmd, "error")
	cmd = fmt.Sprintf("%s && echo %s || echo %s\n", cmd, s.currentCommandOK, s.currentCommandError)
	s.currentCommandOK += "\n"
	s.currentCommandError += "\n"
	fmt.Println("命令", cmd)
	_, err := s.in.Write([]byte(cmd))
	if err != nil {
		if err != io.EOF {
			return "", true, err
		}
	}
	res := <-s.commandChan
	if res != nil {
		return res.data, res.complete, res.err
	}
	return "", true, nil
}

// ReadNext 读取下一条输出
func (s *SSH) ReadNext() (string, bool, error) {
	res := <-s.commandChan
	if res != nil {
		return res.data, res.complete, res.err
	}
	return "", true, nil
}

// Input 输入参数
func (s *SSH) Input(data string) (string, bool, error) {
	_, err := s.in.Write([]byte(data + "\n"))
	if err != nil {
		return "", true, err
	}
	res := <-s.commandChan
	if res != nil {
		return res.data, res.complete, res.err
	}
	return "", true, nil
}

// RunShell 启动交互
func (s *SSH) RunShell(fn func(msg string, err error, errFromCommand bool)) error {
	s.commandChan = make(chan *response, 10)
	s.ch = make(chan *pipObject, 10)
	if err := s.runStdout(s.ch); err != nil {
		return err
	}

	if err := s.runStderr(s.ch); err != nil {
		return err
	}

	var err error
	s.in, err = s.Session.StdinPipe()
	if err != nil {
		return err
	}

	if err := s.Session.Shell(); err != nil {
		// sendErrorLog(rec.ID, params.StepConvert, "执行远程命令 Shell() 时发生错误:"+err.Error())
		return err
	}

	go s.pipProccess(s.ch, fn)

	go func() {
		if err := s.Session.Wait(); err != nil {
			if !IsShellExit(err) {
				fmt.Print(base.ErrorStack(err))
				if fn != nil {
					fn("", err, false)
				}
				s.Close()
				// sendErrorLog(rec.ID, params.StepConvert, "执行远程命令 Wait() 时发生错误:"+err.Error())
				// return !retrying, nil
			}
		}
	}()

	return nil
}

// func (s *SSH) proccess

type pipObject struct {
	byCommand bool
	err       error
	data      string
}

func (s *SSH) pipProccess(ch chan *pipObject, fn func(msg string, err error, errFromCommand bool)) {
	for po := range ch {
		// fmt.Println("po.data", po.data)
		if strings.HasSuffix(po.data, s.currentCommandOK) {
			// 命令执行完成
			po.data = strings.ReplaceAll(po.data, s.currentCommandOK, "")
			s.currentCommand = ""
			s.commandChan <- &response{
				data:     po.data,
				complete: true,
			}
		} else if strings.HasSuffix(po.data, s.currentCommandError) {
			// 命令执行完成

			po.data = strings.ReplaceAll(po.data, s.currentCommandError, "")
			s.currentCommand = ""
			s.commandChan <- &response{
				err:      errors.New(po.data),
				complete: true,
			}
		} else {
			s.commandChan <- &response{
				data:     po.data,
				complete: false,
			}
		}
		if fn != nil {
			fn(po.data, po.err, po.byCommand)
		}
	}
}

func (s *SSH) proccessOutput(out io.Reader, ch chan *pipObject, fn func(buf []byte, dataSize int) *pipObject) {
	buf := make([]byte, 4096)
	for {
		sz, err := out.Read(buf)
		if err != nil {
			if io.EOF != err {
				// 发生错误
				ch <- &pipObject{
					byCommand: false,
					err:       err,
				}
			}
			// 结束
			break
		}
		if sz > 0 {
			// 有数据输出
			ch <- fn(buf, sz)
		}
	}
}

// stdout 标准输出
func (s *SSH) runStdout(ch chan *pipObject) error {
	out, err := s.Session.StdoutPipe()
	if err != nil {
		return err
	}
	go s.proccessOutput(out, ch, func(buf []byte, dataSize int) *pipObject {
		return &pipObject{
			byCommand: true,
			data:      string(buf[:dataSize]),
		}
	})
	return nil
}

// stderr 错误输出
func (s *SSH) runStderr(ch chan *pipObject) error {
	out, err := s.Session.StderrPipe()
	if err != nil {
		return err
	}
	go s.proccessOutput(out, ch, func(buf []byte, dataSize int) *pipObject {
		return &pipObject{
			byCommand: true,
			err:       errors.New(string(buf[:dataSize])),
		}
	})
	return nil
}

// Close 关闭连接
func (s *SSH) Close() error {
	if s.ch != nil {
		close(s.ch)
		s.ch = nil
	}
	if s.commandChan != nil {
		close(s.commandChan)
		s.commandChan = nil
	}

	if s.client != nil {
		err := s.client.Close()
		s.client = nil

		return err
	}
	return nil
}
