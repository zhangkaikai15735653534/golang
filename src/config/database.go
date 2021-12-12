package config

import "github.com/spf13/viper"

/**
* @Author: kiki.zang
* @Date: 2021/12/12 9:00 PM
 */
type Database struct {
	Dbtype   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

/**
 * 项目启动时  加载配置文件 初始化database
 */
func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Dbtype:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
		Name:     cfg.GetString("name"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}

var DatabaseConfig = new(Database)
