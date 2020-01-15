package configure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// ServiceConf 对外服务配置信息
type ServiceConf struct {
	Addr         string `json:"addr" yaml:"addr"`
	Port         int    `json:"port" yaml:"port"`
	WebURL       string `json:"webUrl" yaml:"webUrl"`
	WebFilesPath string `json:"webFilesPath" yaml:"webFilesPath"`
	APIDoc       bool   `json:"apiDoc" yaml:"apiDoc"`
	APIURL       string `json:"apiUrl" yaml:"apiUrl"`
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

// CreateHttpServer 根据配置信息创建 http Server
func (s *ServiceConf) CreateHttpServer(apiHandler http.Handler, log *logrus.Logger) (*http.Server, *http.ServeMux) {
	mux := &http.ServeMux{}
	log.Infof("Server is bind to [%s]", s.HttpAddr())
	svr := &http.Server{
		Addr:    s.HttpAddr(),
		Handler: mux,
	}
	log.Infof("Api url bind to [%s]", s.GetApiURL())
	mux.Handle(s.GetApiURL(), apiHandler)
	if s.WebFilesPath != "" {
		log.Infof("Static page bind to [%s]", s.GetWebURL())
		mux.Handle(s.GetWebURL(), http.FileServer(http.Dir(s.WebFilesPath)))
	}

	return svr, mux
}
