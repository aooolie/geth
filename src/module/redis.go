package module

import (
    "conf"
    "fmt"
    "github.com/garyburd/redigo/redis"
)



func RedisWrite(msg map[string]string) {
    redisUrl := conf.SERVER_IP + ":" + conf.SERVER_PORT
    c, err := redis.Dial("tcp", redisUrl)
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()
    for country := range msg {
        _, err = c.Do("SET", country, msg [country])
        if err != nil {
            fmt.Println("redis set failed:", err)
        }
    }
}

func RedisSingleWrite(key string, value string) {
    redisUrl := conf.SERVER_IP + ":" + conf.SERVER_PORT
    c, err := redis.Dial("tcp", redisUrl)
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()

        _, err = c.Do("SET", key, value)
        if err != nil {
            fmt.Println("redis set failed:", err)
        }
}

func RedisRead(key string) string{
    redisUrl := conf.SERVER_IP + ":" + conf.SERVER_PORT
    c, err := redis.Dial("tcp", redisUrl)
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return ""
    }
    defer c.Close()

    result, err := redis.String(c.Do("GET", key))
    if err != nil {
        fmt.Println("redis get failed:", err)
    }
    return result
}

