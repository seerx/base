package configure

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// GetConfigureData 获取运行环境信息，即配置文件
func GetConfigureData(configureFileArgName string, configureEnvArgName string) ([]byte, error) {
	var file string
	var env string
	msgFile := fmt.Sprintf("Please input configure file path with -%s", configureFileArgName)
	msgEnv := fmt.Sprintf("Please input configure environment name with -%s", configureEnvArgName)
	flag.StringVar(&file, configureFileArgName, "", msgFile)
	flag.StringVar(&env, configureEnvArgName, "", msgEnv)
	flag.Parse()
	var errOfFile error
	if file != "" {
		// 配置文件优先
		fd, err := os.Open(file)
		if err != nil {
			errOfFile = fmt.Errorf("open configure file error: [%w]", err)
			goto second
			// return nil, fmt.Errorf("Open configure file error: [%w]", err)
		}
		defer fd.Close()
		val, err := ioutil.ReadAll(fd)
		if err != nil {
			errOfFile = fmt.Errorf("read configure file error: [%w]", err)
			goto second
			// return nil, fmt.Errorf("Read configure file error: [%w]", err)
		}
		// 返回文件内容
		return val, nil
	}
second:
	if env != "" {
		// 第二选择,使用环境变量，此时可能同时返回错误信息，该错误只需要记录即可
		// 第一个参数不是 nil
		cfgEnv := os.Getenv(env)
		if cfgEnv == "" {
			// 环境变量中没有对应的变量
			return nil, fmt.Errorf("no environment item set by name [%s]", env)
		}
		return []byte(cfgEnv), errOfFile
	}
	if errOfFile != nil {
		// 没有环境变量，但是读取配置文件时发生错误，返回错误信息
		return nil, errOfFile
	}
	return nil, fmt.Errorf("please input configurations from file by -%s or from environment by -%s", configureFileArgName, configureEnvArgName)
}
