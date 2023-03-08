package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// 数据库和缓存配置
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	// user和product的RPC客户端配置
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
}
