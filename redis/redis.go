package redis

import (
	"dario.cat/mergo"
	redisv9 "github.com/redis/go-redis/v9"
	"runtime"
	"time"
)

type (
	// Config 配置
	Config struct {
		// 连接类型，tcp or unix，默认tcp
		Network string `json:"network"`
		// 连接地址
		Addr string `json:"addr"`
		// 客户端名称 每个连接会执行CLIENT SETNAME ClientName命令
		ClientName string `json:"clientName"`
		// 连接用户名
		Username string `json:"username"`
		// 连接密码
		Password string `json:"password"`
		// 数据库实例
		DB int `json:"db"`
		// 最大重试次数 默认3次 -1（不是0!）表示禁用重试
		MaxRetries int `json:"maxRetries"`
		// 最小重试间隔 默认8毫秒 -1（不是0!）表示禁用重试
		MinRetryBackoff time.Duration `json:"minRetryBackoff"`
		// 最大重试间隔 默认512毫秒 -1（不是0!）表示禁用重试
		MaxRetryBackoff time.Duration `json:"maxRetryBackoff"`
		// 建立连接时超时时间 默认5s
		DialTimeout time.Duration `json:"dialTimeout"`
		// 读取超时时间 建议使用超时时间，不要使用阻塞 0: 默认的超时时间 -1: 没有超时时间 -2: 禁用 SetReadDeadline 调用
		ReadTimeout time.Duration `json:"readTimeout"`
		// 写入超时时间 建议使用超时时间，不要使用阻塞 0: 默认的超时时间 -1: 没有超时时间 -2: 禁用 SetWriteDeadline 调用
		WriteTimeout time.Duration `json:"writeTimeout"`
		// 控制客户端是否遵守上下文超时和截止日期 true：遵守 false：不遵守
		ContextTimeoutEnabled bool `json:"contextTimeoutEnabled"`
		// 连接池类型 true: FIFO池() false: LIFO池
		PoolFIFO bool `json:"poolFIFO"`
		// 连接池大小 默认为 runtime.GOMAXPROCS*10 个连接
		PoolSize int `json:"poolSize"`
		// 当连接池中的所有连接都在使用时，客户端等待连接的时间，默认为ReadTimeout + 1秒，如果超时，则返回错误
		PoolTimeout time.Duration `json:"poolTimeout"`
		// 最小空闲连接数，新连接比较慢。 默认为0， 空闲连接不会被关闭
		MinIdleConns int `json:"minIdleConns"`
		// 最大的空闲连接数 默认为0，空闲连接不会被关闭
		MaxIdleConns int `json:"maxIdleConns"`
		// 最大的活跃连接数 默认为0，活跃连接不会被关闭
		MaxActiveConns int `json:"maxActiveConns"`
		// 连接最大空闲时间，应该小于服务器的超时时间，过期的连接可能会在重用之前被懒惰地关闭，如果d<=0，则不会关闭连接。连接的空闲时间，默认为30分钟，-1禁用空闲超时检查
		ConnMaxIdleTime time.Duration `json:"connMaxIdleTime"`
		// 一个连接可以被重用的最大时间，过期的连接可能会在重用之前被懒惰地关闭，如果d<=0，则不会关闭连接。连接的最大生存时间，默认为0，-1禁用连接的最大生存时间
		ConnMaxLifetime time.Duration `json:"connMaxLifetime"`
		// Hook中间件
		Hooks []redisv9.Hook
	}
)

func NewClient(config *Config) *redisv9.Client {
	_ = mergo.Merge(config, defaultConfig())

	rdb := redisv9.NewClient(&redisv9.Options{
		Network:               config.Network,
		Addr:                  config.Addr,
		ClientName:            config.ClientName,
		Username:              config.Username,
		Password:              config.Password,
		DB:                    config.DB,
		MaxRetries:            config.MaxRetries,
		MinRetryBackoff:       config.MinRetryBackoff,
		MaxRetryBackoff:       config.MaxRetryBackoff,
		DialTimeout:           config.DialTimeout,
		ReadTimeout:           config.ReadTimeout,
		WriteTimeout:          config.WriteTimeout,
		ContextTimeoutEnabled: config.ContextTimeoutEnabled,
		PoolFIFO:              config.PoolFIFO,
		PoolSize:              config.PoolSize,
		PoolTimeout:           config.PoolTimeout,
		MinIdleConns:          config.MinIdleConns,
		MaxIdleConns:          config.MaxIdleConns,
		MaxActiveConns:        config.MaxActiveConns,
		ConnMaxIdleTime:       config.ConnMaxIdleTime,
		ConnMaxLifetime:       config.ConnMaxLifetime,
	})

	return rdb
}

func defaultConfig() *Config {
	return &Config{
		Network:               "tcp",
		Addr:                  "localhost:6379",
		DB:                    0,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           3 * time.Second,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              runtime.NumCPU() * 10,
		PoolTimeout:           5 * time.Second,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		MaxActiveConns:        0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
	}
}
