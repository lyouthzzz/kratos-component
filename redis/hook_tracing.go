package redis

import redisv9 "github.com/redis/go-redis/v9"

var _ redisv9.Hook = (*tracingHook)(nil)

func TracingHook() redisv9.Hook { return &tracingHook{} }

type tracingHook struct{}

func (hook *tracingHook) DialHook(next redisv9.DialHook) redisv9.DialHook {
	//TODO implement me
	panic("implement me")
}

func (hook *tracingHook) ProcessHook(next redisv9.ProcessHook) redisv9.ProcessHook {
	//TODO implement me
	panic("implement me")
}

func (hook *tracingHook) ProcessPipelineHook(next redisv9.ProcessPipelineHook) redisv9.ProcessPipelineHook {
	//TODO implement me
	panic("implement me")
}
