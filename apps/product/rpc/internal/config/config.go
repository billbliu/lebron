package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	// Model缓存配置
	CacheRedis cache.CacheConf
	// 业务缓存配置
	BizRedis redis.RedisConf
}
