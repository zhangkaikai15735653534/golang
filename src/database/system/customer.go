package system

import (
	"GoMvc/src/database"
	"fmt"
)

/**
* 用户表
* @Author: kiki.zang
* @Date: 2021/12/13 9:21 PM
 */

/**
 * 定义表名称及结构 在项目启动时进行初始化
 */
type SysCustomer struct {
	SysUserId
	LoginM
	SysUserB
}

type SysUserId struct {
	//定义主键
	CustomerId int `gorm:"primary_key;AUTO_INCREMENT  COMMENT '用户id'"  json:"customerId"` // 用户id
}

type SysUserB struct {
	//todo : 先自己手写一套 后期搞个自动生成
	CustomerAccount  string `gorm:"type:varchar(20)  COMMENT '用户手机'" json:"customerAccount"` //用户手机
	CustomerPassword string `gorm:"type:varchar(200)  COMMENT '是否删除标志'" json:"customerAccount"`
	CustomerAaltVal  string `gorm:"type:varchar(200)   COMMENT '盐'" json:"customerAaltVal"`
	DelFlag          int    `gorm:"type:int(4)   COMMENT '是否删除标志 0：否，1：是'" json:"delFlag"`
	LoginErrorTime   int    `gorm:"type:int(11)  COMMENT '登录错误次数'" json:"loginErrorTime"`
	OpenId           string `gorm:"type:varchar(11)  COMMENT '小程序用户openid'" json:"openId"`
	Avatar           string `gorm:"type:varchar(255)  COMMENT '头像'" json:"avatar"` //头像
	Sex              string `gorm:"type:varchar(255)  COMMENT '性别'" json:"sex"`    //性别
	Email            string `gorm:"type:varchar(128)  COMMENT '邮箱'" json:"email"`  //邮箱
	// 这个不知道做啥的
	Params string `gorm:"-" json:"params"`
	//共有model
	BaseModel
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
}

/**
 * 用户信息
 */
type LoginM struct {
	UserName
}

//声明表名称
func (SysCustomer) TableName() string {
	return "sys_customer"
}

//查询用户信息  *指向SysCustomer  可以直接操作数据库
func (user *SysCustomer) GetUserInfo() (userInfo SysCustomer, err error) {
	table := database.Eloquent.Table(user.TableName()).Select([]string{"sys_customer.*"})
	//todo :这样封装不行   稍后看看这么做  grom
	//if user.Username != "" {
	//	fmt.Println("username is ", user.Username)
	//	table.Where("username = ?", user.Username)
	//}
	//
	//if user.CustomerPassword != "" {
	//	fmt.Println("customer_password is ", user.CustomerPassword)
	//	table.Where("customer_password  = ? ", user.CustomerPassword)
	//}
	//执行查询   &userInfo是正确时返回的

	r := table.Select(&userInfo)
	fmt.Println(r.Row())
	if err = table.Where("username = ?", user.Username).Where("customer_password  = ? ", user.CustomerPassword).Find(&userInfo).Error; err != nil {
		fmt.Println("error is not null")
		return
	}

	return
}
