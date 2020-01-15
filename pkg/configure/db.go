package configure

import "fmt"

type DBConf struct {
	Type           string `yaml:"type"`           // 数据库: pgsql,mysql ...
	Host           string `yaml:"host"`           // 数据库地址
	Port           int    `yaml:"port"`           // 端口
	User           string `yaml:"user"`           // 用户
	Passwd         string `yaml:"passwd"`         // 密码
	DB             string `yaml:"db"`             // 数据库名称
	SSLMode        string `yaml:"sslMode"`        // pgsql 专用，默认为 disable
	ConnectTimeout int    `yaml:"connectTimeout"` // 连接超时设置，默认为 10
}

// String 生成数据库连接串
func (dbc *DBConf) String() string {
	if dbc.Type == "pgsql" || dbc.Type == "postgresql" {
		return pgsql(dbc)
	}
	panic(fmt.Errorf("Do not support database [%s]", dbc.Type))
}

func pgsql(dbc *DBConf) string {
	const pgTemplate = "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d"
	if dbc.SSLMode == "" {
		dbc.SSLMode = "disable"
	}
	if dbc.ConnectTimeout == 0 {
		dbc.ConnectTimeout = 10
	}
	return fmt.Sprintf(pgTemplate, dbc.Host, dbc.Port, dbc.User, dbc.DB, dbc.Passwd, dbc.SSLMode, dbc.ConnectTimeout)
}
