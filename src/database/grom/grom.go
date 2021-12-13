package grom

import (
	"GoMvc/src/database/system"
	"github.com/jinzhu/gorm"
)

/**
* 数据库迁移  执行数据库初始化操作
* @Author: kiki.zang
* @Date: 2021/12/13 9:18 PM
 */
func AutoMigrate(db *gorm.DB) error {
	//gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中：创建表、添加数据时都是如此。指定了db.SingularTable(true)之后，进行严格匹配
	db.SingularTable(true)
	return db.AutoMigrate(

		//用户信息表初始化
		new(system.SysCustomer),
	).Error
}
