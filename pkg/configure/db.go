package configure

import "fmt"

type DBConf struct {
	Dialect        string `json:"dialect" yaml:"type"`                  // 数据库: postgres,mysql ...
	Host           string `json:"host" yaml:"host"`                     // 数据库地址
	Port           int    `json:"port" yaml:"port"`                     // 端口
	User           string `json:"user" yaml:"user"`                     // 用户
	Password       string `json:"passwd" yaml:"passwd"`                 // 密码
	DB             string `json:"db" yaml:"db"`                         // 数据库名称
	SSLMode        string `json:"sslMode" yaml:"sslMode"`               // pgsql 专用，默认为 disable
	ConnectTimeout int    `json:"connectTimeout" yaml:"connectTimeout"` // 连接超时设置，默认为 10
}

// String 生成数据库连接串
func (dbc *DBConf) String() string {
	if dbc.Dialect == "pgsql" || dbc.Dialect == "postgres" || dbc.Dialect == "postgresql" {
		dbc.Dialect = "postgres"
		return pgsql(dbc)
	}
	panic(fmt.Errorf("Do not support database [%s]", dbc.Dialect))
}

func pgsql(dbc *DBConf) string {
	const pgTemplate = "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d"
	if dbc.SSLMode == "" {
		dbc.SSLMode = "disable"
	}
	if dbc.ConnectTimeout == 0 {
		dbc.ConnectTimeout = 10
	}
	return fmt.Sprintf(pgTemplate, dbc.Host, dbc.Port, dbc.User, dbc.DB, dbc.Password, dbc.SSLMode, dbc.ConnectTimeout)
}
