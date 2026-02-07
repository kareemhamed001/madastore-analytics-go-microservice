package utils

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

var errCacheNotInitialized = errors.New("cache not initialized")

func InitCache(addr, password string, db int, dialTimeout, readTimeout, writeTimeout time.Duration) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rdb.Ping(ctx).Err()
}

func CloseCache() error {
	if rdb == nil {
		return nil
	}
	return rdb.Close()
}

func SetCache(key string, value interface{}, ttl time.Duration) error {
	if rdb == nil {
		return errCacheNotInitialized
	}
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return rdb.Set(ctx, key, b, ttl).Err()
}

func GetCache(key string, dest interface{}) (bool, error) {
	if rdb == nil {
		return false, errCacheNotInitialized
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, json.Unmarshal([]byte(val), dest)
}
