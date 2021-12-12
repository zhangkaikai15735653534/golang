package handler

import "github.com/gin-gonic/gin"

/**
* 默认响应类
* @Author: kiki.znag
* @Date: 2021/12/12 4:01 PM
 */
func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
