package init_project

import (
	"init_project/internal/database"
	"init_project/internal/utils"
)

func InitProject()  error{
	//初始化日志模块
	logger := utils.Logger
	err := logger.Init(utils.Config.LogConfig.RuntimeConf, utils.Config.LogConfig.WorkConf)
	if err != nil {
		return err
	}

	//初始化redis
	if err := database.InitRedis(); err != nil {
		return err
	}
	//defer database.CloseRedis()

	//初始化db
	database.InitDB()

	//初始化aerosplike
	_, err = database.InitAerosplike(utils.Config.AerospikeConfig.AerospikeAddr, utils.Config.AerospikeConfig.AerospikePort)
	if err != nil {
		return err
	}
	return nil
}
