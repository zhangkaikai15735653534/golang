package database

import "GoMvc/src/config"

/**
* @Author: kiki.zang
* @Date: 2021/12/12 9:10 PM
 */
func Setup() {
	dbType := config.DatabaseConfig.Dbtype
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}
}
