package config

type DbConfig struct {
	// 数据库用户名
	Username string
	// 密码
	Password string
	Host string
	Port string
	// 数据库名
	DbName string
	// 最大闲置连接
	MaxIdleConns int
	// 最大打开连接
	MaxOpenConns int
}


func NewDbConfig() *DbConfig {
	pwd := "root"
	uname := "root"
	dbname := "homework"
	conf := &DbConfig{Username:uname, Password: pwd, Host: "localhost", Port: "3306", DbName: dbname, MaxIdleConns: 5, MaxOpenConns: 100}
	return conf
}