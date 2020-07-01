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
	Host         string `json:"host" yaml:"host"`
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

// GetWebURL 获取 web 地址
func (s *ServiceConf) GetWebURL() string {
	return checkPath(s.WebURL)
}

// GetAPIURL 获取 api 地址
func (s *ServiceConf) GetAPIURL() string {
	return checkPath(s.APIURL)
}

// HTTPAddr 生成 http 监听的地址
func (s *ServiceConf) HTTPAddr() string {
	return fmt.Sprintf("%s:%d", s.Addr, s.Port)
}

// CreateHTTPServer 根据配置信息创建 http Server
func (s *ServiceConf) CreateHTTPServer(apiHandler http.Handler, log *logrus.Logger) (*http.Server, *http.ServeMux) {
	mux := &http.ServeMux{}
	log.WithField("addr", s.HTTPAddr()).Infof("Server is listening")
	svr := &http.Server{
		Addr:    s.HTTPAddr(),
		Handler: mux,
	}
	log.WithField("api", fmt.Sprintf("%s%s", s.HTTPAddr(), s.GetAPIURL())).Infof("API")
	mux.Handle(s.GetAPIURL(), apiHandler)
	if s.WebFilesPath != "" {
		log.WithField("api", fmt.Sprintf("%s%s", s.GetWebURL(), s.GetAPIURL())).Infof("Static page")
		mux.Handle(s.GetWebURL(), http.FileServer(http.Dir(s.WebFilesPath)))
	}

	return svr, mux
}
