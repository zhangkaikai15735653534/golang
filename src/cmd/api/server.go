package api

import (
	comfig2 "GoMvc/src/config"
	"GoMvc/src/database"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

/**
* @Author: kiki.zang
* @Date: 2021/12/12 9:28 PM
 */
var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		//Use:     "server",
		//Short:   "Start API server",
		//Example: "ferry server config/setting.yml",
		//PreRun: func(cmd *cobra.Command, args []string) {
		//	fmt.Println("gogogo  init")
		//	//usage() //项目启动
		//	//setup() //加载配置项
		//},
		//RunE: func(cmd *cobra.Command, args []string) error {
		//	return run()
		//},
	}
)

/**
 * 初始化配置文件
 */
func init() {
	//todo : 路径有问题  稍后修复
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "/Users/kiki/workspace/go/GoMvc/src/config/setting.yml", "Start server with provided configuration file")
	//StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	//StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func Usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func Setup() {
	// 1. 读取配置文件setting.yml
	comfig2.ConfigSetup(config)
	// 2. 初始化数据库链接
	database.Setup()
	fmt.Println("init success")
}
