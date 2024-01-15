package redis

import redisv9 "github.com/redis/go-redis/v9"

var _ redisv9.Hook = (*metricsHook)(nil)

func MetricsHook() redisv9.Hook { return &metricsHook{} }

type metricsHook struct{}

func (hook *metricsHook) DialHook(next redisv9.DialHook) redisv9.DialHook {
	//TODO implement me
	panic("implement me")
}

func (hook *metricsHook) ProcessHook(next redisv9.ProcessHook) redisv9.ProcessHook {
	//TODO implement me
	panic("implement me")
}

func (hook *metricsHook) ProcessPipelineHook(next redisv9.ProcessPipelineHook) redisv9.ProcessPipelineHook {
	//TODO implement me
	panic("implement me")
}
