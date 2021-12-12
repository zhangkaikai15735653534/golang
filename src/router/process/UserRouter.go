package process

import (
	"GoMvc/src/controller"
	"github.com/gin-gonic/gin"
)

/**
* @Author: kiki.zang
* @Date: 2021/12/12 7:42 PM
 */

/**
 * 用户路由
 */
func UserRouter(v1 *gin.RouterGroup) {
	//路由分发 并执行jwt校验
	userGroup := v1.Group("/getUser").Use(controller.CheckJwt())
	{
		//校验通过了  走下边的流程吧  必须换一行 否则报错了 go这个规则 啧啧
		//获取用户信息
		userGroup.POST("/userinfo", controller.LoginAfter)
	}

}
