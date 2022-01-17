package database

import (
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/quexer/red"
	"github.com/quexer/utee"
)

var (
	redOnce  sync.Once
	redisObj *Redis
)

// RedisOption Option redis连接选项
type RedisOption struct {
	Address  string
	Password string
	Db       int // db number
}

// Redis 包装对象
type Redis struct {
	Pool *redis.Pool
	Do   red.DoFunc
}

// NewRedis 初始化redis连接，返回包装对象
func NewRedis(opt RedisOption) *Redis {
	redOnce.Do(func() {
		pool := red.CreatePool(600, opt.Address, opt.Password, opt.Db)
		redisDo := red.BuildDoFunc(pool)
		_, err := redisDo("PING")
		utee.Chk(err)

		redisObj = &Redis{
			Pool: pool,
			Do:   redisDo,
		}
	})

	return redisObj
}

// SaveInt 保存数值 expire(秒)
func (p *Redis) SaveInt(key string, val int, expire int) error {
	if expire > 0 {
		_, err := p.Do("SETEX", key, expire, val)
		return err
	}
	_, err := p.Do("SET", key, val)
	return err
}

// GetInt 按键获取数值
func (p *Redis) GetInt(key string) (int, error) {
	i, err := redis.Int(p.Do("GET", key))

	if err != nil && err == redis.ErrNil {
		// expire
		return 0, nil
	}
	return i, err
}

// MustGetInt 按键一定会获取数值，否则返回错误
func (p *Redis) MustGetInt(key string) (int, error) {
	return redis.Int(p.Do("GET", key))
}

// SaveString 保存字符串
func (p *Redis) SaveString(key string, val string, expire int) error {
	if expire > 0 {
		_, err := p.Do("SETEX", key, expire, val)
		return err
	}
	_, err := p.Do("SET", key, val)
	return err

}

// GetString 按键获取到数据
func (p *Redis) GetString(key string) (string, error) {
	s, err := redis.String(p.Do("GET", key))

	if err != nil && err == redis.ErrNil {
		return "", nil
	}
	return s, err
}

// MustGetString 按键一定会获取到字符串数据，否则返回错误
func (p *Redis) MustGetString(key string) (string, error) {
	return redis.String(p.Do("GET", key))
}

// TTL 检查过期时间， 返回秒
func (p *Redis) TTL(key string) (int, error) {
	n, err := redis.Int(p.Do("TTL", key))

	if err != nil && err == redis.ErrNil {
		// expire
		return 0, nil
	}
	return n, err
}

// DelKey 按键删除
func (p *Redis) DelKey(key string) error {
	_, err := p.Do("DEL", key)
	return err
}

func (p *Redis) AutoCacheString(key string, useCache bool, f func() (string, int, error)) (string, error) {
	var result string

	if useCache {
		s, err := p.GetString(key)
		if err != nil {
			return "", err
		}
		result = s
	}

	if result == "" {
		s, expire, err := f()
		if err != nil {
			return "", err
		}
		err = p.SaveString(key, s, expire)
		if err != nil {
			return "", err
		}
		result = s
	}
	return result, nil
}
