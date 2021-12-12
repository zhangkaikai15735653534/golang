package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
* @Author: kiki.znag
* @Date: 2021/12/12 4:23 PM
 */

/**
 * 获取验证码
 */
func GenerateCaptchaHandler(c *gin.Context) {
	fmt.Println("get code begin====")
}

func Login(c *gin.Context) {
	//定义变量接收返回值
	var (
		res bool
		err error
	)
	fmt.Println("login  begin")
	name := c.PostForm("userName")
	pwd := c.PostForm("password")
	fmt.Println(name)
	fmt.Println(pwd)
	res, err = checkUser(name, pwd)
	if res {
		fmt.Println("校验通过 ，开始执行下一步操作")
		//todo:查询数据库获取用户信息
		var userId int64 = 100
		var userName string = "kiki"
		//生成jwt 并返回
		c.JSON(200, gin.H{
			"message": "ok",
			"token":   GetJwt(userId, userName),
		})

	} else {
		//校验未通过  直接返回
		fmt.Println("校验未通过  return，", err)
		c.JSON(-1, gin.H{
			"message": "用户名或密码错误",
		})
	}

}

func checkUser(name string, pwd string) (bool, error) {
	var err error
	if name == "" {
		fmt.Println("this name is null")
		return false, err
	}
	if pwd == "" {
		fmt.Println("this pwd is null")
		return false, err
	}

	if "zkk" == name && "123456" == pwd {
		fmt.Println("check is password")
		return true, nil
	}
	//名字或密码错误
	return false, err
}

/**
* 登陆校验之后的操作
 */
func LoginAfter(c *gin.Context) {
	fmt.Println("登陆校验之后的操作==begin")
}
