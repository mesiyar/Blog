package util

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
)

type Redis struct {
}

var RedisConn *redis.Pool

func (r *Redis) Init() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisMaxIdle,
		MaxActive:   setting.RedisMaxActive,
		IdleTimeout: setting.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisHost)
			if err != nil {
				logging.Warn("redis redis 链接失败!")
				return nil, err
			}
			if setting.RedisPassword != "" {
				if _, err := c.Do("AUTH", setting.RedisPassword); err != nil {
					c.Close()
					logging.Warn("redis 认证失败!")
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func (r *Redis) Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func (r *Redis) Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *Redis) Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func (r *Redis) LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = r.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Redis) HSet(key string, sub string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", key, sub, value)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) HDel(key, sub string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, sub)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) HGet(key string, sub string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", key, sub))
	if err != nil {
		return nil, err
	}

	return data, err
}
