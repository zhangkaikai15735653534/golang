package router

import (
	"GoMvc/src/controller"
	"GoMvc/src/handler"
	"GoMvc/src/router/process"
	"github.com/gin-gonic/gin"
)

/**
* @Author: kiki.znag
* @Date: 2021/12/12 3:58 PM
 */
func InitRouter() *gin.Engine {
	//new gin

	r := gin.New()
	//RouterGroup
	g := r.Group("")
	//默认
	SysBaseRouter(g)

	//非鉴权接口
	NoAuthRouter(g)

	//鉴权接口
	AuthRouter(g)

	return r

}

/**
* @Author: kiki.znag
* @Date: 2021/12/12 4:00 PM
 */

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/info", handler.Info)
}

/**
 * 非鉴权路由
 */
func NoAuthRouter(g *gin.RouterGroup) {
	v1 := g.Group("/api/v1")
	//
	v1.GET("/getCaptcha", controller.GenerateCaptchaHandler)

}

/**
 * 鉴权路由
 */
func AuthRouter(g *gin.RouterGroup) {
	//在执行以下路由前
	g.POST("/login", controller.Login)

	//路由分发规则 感觉相当于  requestmapping
	v1 := g.Group("/api/v2") //api/v2路由走这里  use是鉴权接口

	//优化一下 区分一下模块

	//用户路由  /getUser  总路由(自己起的名字)
	process.UserRouter(v1)

}
