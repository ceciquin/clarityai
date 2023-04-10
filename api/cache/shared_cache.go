package cache

import (
    "time"

    "github.com/go-redis/redis"
)

type Cache struct {
    client *redis.Client
}

func NewCache(addr string, password string, db int) (*Cache, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    _, err := client.Ping().Result()
    if err != nil {
        return nil, err
    }

    return &Cache{client}, nil
}

func (c *Cache) Get(key string) (string, error) {
    val, err := c.client.Get(key).Result()
    if err == redis.Nil {
        return "", nil
    } else if err != nil {
        return "", err
    }

    return val, nil
}

func (c *Cache) Set(key string, value string, expiration time.Duration) error {
    err := c.client.Set(key, value, expiration).Err()
    if err != nil {
        return err
    }

    return nil
}

func (c *Cache) Delete(key string) error {
    err := c.client.Del(key).Err()
    if err != nil {
        return err
    }

    return nil
}
