package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// jwt验证
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	// 数据库和缓存在rpc端操纵
	// 只需要配置RPC client
	ProductRpc zrpc.RpcClientConf
}
