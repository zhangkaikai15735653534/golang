package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

/**
* @Author: kiki.zang
* @Date: 2021/12/12 9:40 PM
 */

var cfgDatabase *viper.Viper

//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error 1")
		//logger.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println("error 2")
		//logger.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	// 数据库初始化  读取setting下的database
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	//初始化数据库
	DatabaseConfig = InitDatabase(cfgDatabase)
}
