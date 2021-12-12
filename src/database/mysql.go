package database

import (
	"GoMvc/src/config"
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"strconv"
	//todo: 默认加载mysql配置
	_ "github.com/go-sql-driver/mysql" //加载mysql
)

/**
* @Author: kiki.zang
* @Date: 2021/12/12 8:59 PM
 */

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func (e *Mysql) Setup() {

	var err error
	var db Database

	db = new(Mysql)
	//获取数据库连接
	MysqlConn = db.GetConnect()

	Eloquent, err = db.Open(DbType, MysqlConn)

	if err != nil {
		fmt.Println("%s connect error %v", DbType, err)
	} else {
		fmt.Println("%s connect success!", DbType)
	}

	if Eloquent.Error != nil {
		fmt.Println("database error %v", Eloquent.Error)
	}

	// 是否开启详细日志记录
	Eloquent.LogMode(viper.GetBool("settings.gorm.logMode"))

	// 设置最大打开连接数
	Eloquent.DB().SetMaxOpenConns(viper.GetInt("settings.gorm.maxOpenConn"))

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用

	Eloquent.DB().SetMaxIdleConns(viper.GetInt("settings.gorm.maxIdleConn"))
}

type Mysql struct {
}

func (e *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

func (e *Mysql) GetConnect() string {

	//建立数据库连接

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	return conn.String()
}
