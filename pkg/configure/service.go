package configure

import (
	"fmt"
	"strings"
)

// ServiceConf 对外服务配置信息
type ServiceConf struct {
	Addr         string `yaml:"addr"`
	Port         int    `yaml:"port"`
	WebURL       string `yaml:"webUrl"`
	WebFilesPath string `yaml:"webFilesPath"`
	APIDoc       bool   `yaml:"apiDoc"`
	APIURL       string `yaml:"apiUrl"`
}

func checkPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}

// GetApiURL 获取 web 地址
func (s *ServiceConf) GetWebURL() string {
	return checkPath(s.WebURL)
}

// GetApiURL 获取 api 地址
func (s *ServiceConf) GetApiURL() string {
	return checkPath(s.APIURL)
}

// HttpAddr 生成 http 监听的地址
func (s *ServiceConf) HttpAddr() string {
	return fmt.Sprintf("%s:%d", s.Addr, s.Port)
}
