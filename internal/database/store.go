package database

import "init_project/internal/utils"

var Mysql *DataStore
var Gp *DataStore
var GpV2 *DataStore

func InitDB() {
	Mysql = NewDB(MysqlDriver, utils.Config.GetMysqlConfig())
}
