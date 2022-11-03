package main

import (
	"flag"
	"fmt"
	"init_project/internal/database"
	"init_project/internal/service"
	"init_project/internal/utils"
	"init_project/routers"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func main() {
	//设置并发度
	CORE_NUM := runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(CORE_NUM * 4)

	//处理命令行参数
	confPath := flag.String("c", "./conf/config.conf", "configure file path")
	flag.Parse()
	if err := utils.LoadConfigFile(*confPath); err != nil {
		fmt.Println("load config failed: ", err.Error())
		os.Exit(1)
	}

	//初始化日志模块
	logger := utils.Logger
	err := logger.Init(utils.Config.LogConfig.RuntimeConf, utils.Config.LogConfig.WorkConf)
	if err != nil {
		fmt.Println("InitLog failed: ", err.Error())
		os.Exit(1)
	}

	//初始化redis
	if err := database.InitRedis(); err != nil {
		fmt.Println("InitRedis failed: ", err.Error())
		os.Exit(1)
	}
	defer database.CloseRedis()
	//初始化db
	database.InitDB()
	//初始化http服务
	service := service.NewService()
	loadRouters := routers.LoadRouters(service)
	http.ListenAndServe(":"+strconv.Itoa(utils.Config.HttpConfig.Port), loadRouters)
}