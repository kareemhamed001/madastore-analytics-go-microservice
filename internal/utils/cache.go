package utils

import (
    "context"
    "encoding/json"
    "time"

    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

func SetCache(key string, value interface{}, ttl time.Duration) error {
    b, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return rdb.Set(ctx, key, b, ttl).Err()
}

func GetCache(key string, dest interface{}) (bool, error) {
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return false, nil
    } else if err != nil {
        return false, err
    }
    return true, json.Unmarshal([]byte(val), dest)
}
