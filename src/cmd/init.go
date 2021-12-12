package cmd

import (
	"GoMvc/src/router"
)

/**
* @Author: kiki.znag
* @Date: 2021/12/12 3:54 PM
 */

/**
项目初始化时执行
*/
func Execute() {
	//注册系统路由
	r := router.InitRouter()

	//监听端口号
	r.Run(":9999")
}
