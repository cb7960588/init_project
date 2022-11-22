package main

import (
	"flag"
	"fmt"
	"init_project/init_project"
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

	err := init_project.InitProject()
	if err != nil {
		fmt.Println("init project failed: ",err.Error())
		os.Exit(1)
	}

	



	//初始化http服务
	service := service.NewService()
	loadRouters := routers.LoadRouters(service)
	http.ListenAndServe(":"+strconv.Itoa(utils.Config.HttpConfig.Port), loadRouters)
}