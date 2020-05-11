package base

import (
	"os"
	"os/exec"
	"strings"
)

// GetSelfPath 获取可执行程序所在路径
func GetSelfPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	// return s
	i := strings.LastIndex(s, string(os.PathSeparator))
	// fmt.Println(i)
	path := string(s[0 : i+1])
	return path
}

// PathType 路径类型
type PathType int

const (
	// PTNone 未定义
	PTNone PathType = iota
	// PTFile 文件
	PTFile
	// PTDirectory 目录
	PTDirectory
)

// PathExists 路径是否存在
func PathExists(path string) (bool, PathType, error) {
	st, err := os.Lstat(path)
	if err == nil {
		return true, PTNone, nil
	}
	if os.IsNotExist(err) {
		return false, PTNone, nil
	}
	if st.IsDir() {
		return false, PTDirectory, err
	}
	return false, PTFile, err
}

// CheckDirs 检查路径是否存在，不存在则创建路径
func CheckDirs(path string) bool {
	exists, _, err := PathExists(path)
	if err != nil {
		return false
	}

	if !exists {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}

	return true
}