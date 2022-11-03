package database

import (
	"github.com/go-redis/redis"
	"init_project/internal/utils"
	"time"
)

var redisClient *redis.ClusterClient

func InitRedis() error {
	clusterOptions := &redis.ClusterOptions{
		Addrs:              []string{utils.Config.RedisConfig.Addr},
		DialTimeout:        time.Duration(utils.Config.RedisConfig.ConnectTimeout * 1000000),
		ReadTimeout:        time.Duration(utils.Config.RedisConfig.ReadTimeout * 1000000),
		WriteTimeout:       time.Duration(utils.Config.RedisConfig.WriteTimeout * 1000000),
		IdleTimeout:        -1,
		IdleCheckFrequency: -1,
	}
	if utils.Config.RedisConfig.Password != "" {
		clusterOptions.Password = utils.Config.RedisConfig.Password
	}
	redisClient = redis.NewClusterClient(clusterOptions)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(key, value, expiration).Err()
}

func Get(key string) (string, error) {
	resp, err := redisClient.Get(key).Result()
	if err == nil { // 正常取出
		return resp, nil
	}
	if err == redis.Nil { // 无此缓存
		return "", nil
	}
	return resp, err // 取出失败
}

func Lock(key string, value interface{}, expiration time.Duration) (bool, error) {
	return redisClient.SetNX(key, value, expiration).Result()
}

func Del(key string) error {
	return redisClient.Del(key).Err()
}

func LPush(listName, value string) error {
	return redisClient.LPush(listName, value).Err()
}

func LPop(listName string) (string, error) {
	value, err := redisClient.LPop(listName).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func RPop(listName string) (string, error) {
	value, err := redisClient.RPop(listName).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func CloseRedis() {
	redisClient.Close()
}
