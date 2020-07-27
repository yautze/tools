package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	m *redis.Client
	ctx = context.Background()
)

// M - return redis Client(other conneation)
func M() (*redis.Client, error) {
	if _, err := m.Ping(ctx).Result(); err != nil {
		log.WithError(err).Error("redis connect failed")
		return nil, err
	}	
	return m, nil
}

// Con - 
func Con(opt * redis.Options) error{
	if err := con(opt); err != nil{
		return err
	}
	return nil
}

func con(opt *redis.Options) error {
	rdb := redis.NewClient(opt)

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.WithError(err).Error("redis connect failed")
		return err
	}

	m = rdb

	return nil
}
