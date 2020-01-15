package configure

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// GetConfigureData 获取运行环境信息，即配置文件
func GetConfigureData(configureFileArgName string, configureEnvArgName string) []byte {
	var file string
	var env string
	msgFile := fmt.Sprintf("Please input configure file path with -%s", configureFileArgName)
	msgEnv := fmt.Sprintf("Please input configure environment name with -%s", configureEnvArgName)
	flag.StringVar(&file, configureFileArgName, "", msgFile)
	flag.StringVar(&env, configureEnvArgName, "", msgEnv)
	flag.Parse()
	if file != "" {
		// 文件优先
		fd, err := os.Open(file)
		if err != nil {
			log.Fatal(fmt.Errorf("Open configure file error: [%w]", err))
		}
		defer fd.Close()
		val, err := ioutil.ReadAll(fd)
		if err != nil {
			log.Fatal(fmt.Errorf("Read configure file error: [%w]", err))
		}

		return val
	}
	if env != "" {
		// 使用环境变量
		return []byte(env)
	}
	log.Fatal(fmt.Errorf("Please input configurations from file by -%s or from environment by -%s", configureFileArgName, configureEnvArgName))
	return nil
}
