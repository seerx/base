package base

import (
	"os"
	"os/exec"
	"path/filepath"
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
		if st.IsDir() {
			return true, PTDirectory, err
		}
		return true, PTFile, nil
	}
	if os.IsNotExist(err) {
		return false, PTNone, nil
	}

	return false, PTNone, err
}

// CheckDirs 检查路径是否存在，不存在则创建路径
func CheckDirs(path string) error {
	exists, _, err := PathExists(path)
	if err != nil {
		return err
	}

	if !exists {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetParentPathAndMyName 计算上级目录和本身名称
func GetParentPathAndMyName(path string) (string, string) {
	path = filepath.Clean(path)
	ary := strings.Split(path, string(filepath.Separator))
	if ary[0] == "" {
		ary[0] = string(filepath.Separator)
	}
	return filepath.Join(ary[0 : len(ary)-1]...), ary[len(ary)-1]
}
