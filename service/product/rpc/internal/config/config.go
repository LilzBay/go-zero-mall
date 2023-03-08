package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// MySQL数据源配置
	Mysql struct {
		DataSource string
	}
	// Redis配置
	CacheRedis cache.CacheConf
}
