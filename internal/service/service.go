package service

import (
	"github.com/gin-gonic/gin"
	"init_project/internal/database"
	"init_project/internal/utils"
	"net/http"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (*Service) HealthCheck(c *gin.Context) {
	/*
		utils.Logger.Work.Infof("info_time= %d",time.Now().Unix())
		utils.Logger.Work.Errorf("error_time= %d",time.Now().Unix())
		utils.Logger.Work.Warnf("warn_time= %d",time.Now().Unix())
		utils.Logger.Work.Debugf("debug_time= %d",time.Now().Unix())
		utils.Logger.Work.Criticalf("criticalf_time= %d",time.Now().Unix())
	*/

	type HealthCheck struct {
		Mysql string `json:"mysql"`
		Redis string `json:"redis"`
	}
	h := HealthCheck{
		Mysql: "ok",
		Redis: "ok",
	}
	_, err := database.Mysql.DbHealthCheck()
	if err != nil {
		h.Mysql = "failed"
		utils.Logger.Work.Errorf("HealthCheck mysql err: %s", err)
	}
	redisHealthCheck := "redisHealthCheck"
	err = database.Set(redisHealthCheck, "redisHealthCheck", time.Second*60)
	if err != nil {
		h.Redis = "failed"
		utils.Logger.Work.Errorf("HealthCheck redis set err: %s", err)
	}
	redisVal, err := database.Get(redisHealthCheck)
	if redisVal != redisHealthCheck || err != nil {
		h.Redis = "failed"
		utils.Logger.Work.Error("HealthCheck redis get err")
	}
	c.JSON(http.StatusOK, h)
}
