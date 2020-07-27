package redis

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_RedisConnect(t *testing.T) {
	c := &redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10, // 最大連線數
		MaxRetries:   3,  // 最大重新嘗試連線次數
		MinIdleConns: 5,  // 最小空閒的連線數(建立新連線是較慢的)
	}

	err := Con(c)
	assert.NoError(t, err)

	redisClient, _ := M()

	res := redisClient.Set(ctx,"TestKey", "TestValue", 0)
	assert.Equal(t, "OK",res.Val())
}
