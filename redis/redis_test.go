package redis

import (
	"context"
	redisv9 "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	t.Run("初始化实例", func(t *testing.T) {
		rdb := NewClient(&Config{
			ClientName:            "localhost",
			Password:              "root",
			DB:                    0,
			MaxRetries:            0,
			MinRetryBackoff:       0,
			MaxRetryBackoff:       0,
			DialTimeout:           0,
			ReadTimeout:           0,
			WriteTimeout:          0,
			ContextTimeoutEnabled: false,
			PoolFIFO:              false,
			PoolSize:              0,
			PoolTimeout:           0,
			MinIdleConns:          0,
			MaxIdleConns:          0,
			MaxActiveConns:        0,
			ConnMaxIdleTime:       0,
			ConnMaxLifetime:       0,
			Hooks:                 nil,
		})
		require.NotNil(t, rdb)
	})

	t.Run("测试日志中间件", func(t *testing.T) {
		rdb := NewClient(&Config{
			Addr:                  "localhost:6379",
			ClientName:            "localhost",
			Username:              "",
			Password:              "root",
			DB:                    0,
			MaxRetries:            0,
			MinRetryBackoff:       0,
			MaxRetryBackoff:       0,
			DialTimeout:           0,
			ReadTimeout:           0,
			WriteTimeout:          0,
			ContextTimeoutEnabled: false,
			PoolFIFO:              false,
			PoolSize:              0,
			PoolTimeout:           0,
			MinIdleConns:          0,
			MaxIdleConns:          0,
			MaxActiveConns:        0,
			ConnMaxIdleTime:       0,
			ConnMaxLifetime:       0,
			Hooks:                 nil,
		})
		rdb.AddHook(LoggerHook())

		require.Nil(t, rdb.Set(context.Background(), "key", "value", 5*time.Second).Err())
		require.Nil(t, rdb.Get(context.Background(), "key").Err())

		_, _ = rdb.Pipelined(context.Background(), func(p redisv9.Pipeliner) error {
			p.Ping(context.Background())
			p.Ping(context.Background())
			return nil
		})

	})

}
