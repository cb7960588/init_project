package service

import (
	as "github.com/aerospike/aerospike-client-go"
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
		As    string `json:"as"`
	}
	h := HealthCheck{
		Mysql: "ok",
		Redis: "ok",
		As:    "ok",
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
	//as set 测试
	askey, _ := as.NewKey(utils.Config.AerospikeConfig.AerospikeNamespace, "testSetName", "keyName")
	testValue := []byte("test_value")
	bin1 := as.NewBin("key", testValue)
	policy := as.NewWritePolicy(0, 86400)
	// Write a record
	err = database.AerospikeClient.PutBins(policy, askey, bin1)
	if err != nil {
		h.As = "failed"
		utils.Logger.Work.Error("HealthCheck as set err")
	}

	//as get 测试
	record, err := database.AerospikeClient.Get(nil, askey, "key")
	if err != nil {
		h.As = "failed"
		utils.Logger.Work.Error("HealthCheck as get err")
	}
	_, ok := record.Bins["key"].([]byte)
	if !ok {
		h.As = "failed"
		utils.Logger.Work.Error("HealthCheck as get err")
	}
	c.JSON(http.StatusOK, h)
}
