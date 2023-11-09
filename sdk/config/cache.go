package config

import (
	"github.com/welleren/go-admin-core/storage"
	"github.com/welleren/go-admin-core/storage/cache"
)

type Cache struct {
	Redis     *RedisConnectOptions
	StatRedis *RedisConnectOptions
	Memory    interface{}
}

// CacheConfig cache配置
var CacheConfig = new(Cache)

// Setup 构造cache 顺序 redis > 其他 > memory
func (e Cache) Setup() (storage.AdapterCache, error) {
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		r, err := cache.NewRedis(GetRedisClient(), options)
		if err != nil {
			return nil, err
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		return r, nil
	}
	return cache.NewMemory(), nil
}

// Setup 构造cache 顺序 redis > 其他 > memory
func (e Cache) SetupStat() (storage.AdapterCache, error) {
	if e.StatRedis != nil {
		options, err := e.StatRedis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		r, err := cache.NewRedis(GetStatRedisClient(), options)
		if err != nil {
			return nil, err
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		return r, nil
	}
	return cache.NewMemory(), nil
}
