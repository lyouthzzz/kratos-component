package redis

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/extra/rediscmd/v9"
	redisv9 "github.com/redis/go-redis/v9"
	"net"
	"runtime"
	"strings"
)

var _ redisv9.Hook = (*loggerHook)(nil)

func LoggerHook() redisv9.Hook { return &loggerHook{} }

type loggerHook struct{}

func (logger *loggerHook) DialHook(next redisv9.DialHook) redisv9.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := next(ctx, network, addr)
		if err != nil {
			log.Context(ctx).Errorf("dial redis client failed. %s %s", network, addr)
		} else {
			log.Context(ctx).Infof("dial redis client success. %s %s", network, addr)
		}
		return conn, err
	}
}

func (logger *loggerHook) ProcessHook(next redisv9.ProcessHook) redisv9.ProcessHook {
	return func(ctx context.Context, cmd redisv9.Cmder) error {
		var (
			kvs   = make([]any, 0)
			level = log.LevelInfo
			err   error
		)

		if err = next(ctx, cmd); redisCmderError(err) != nil {
			level = log.LevelError
			kvs = append(kvs, "error", err.Error())
		}

		fn, file, line := funcFileLine("github.com/redis/go-redis")
		kvs = append(kvs, "msg", fmt.Sprintf("filepath: %s. lineno: %d. function: %s", file, line, fn))

		request := rediscmd.CmdString(cmd)
		reply := rediscmd.Bytes(redisCmderReply(cmd))
		// 请求参数 CMD命令
		kvs = append(kvs, "request", rediscmd.CmdString(cmd))
		// 出口带宽 发送请求
		kvs = append(kvs, "egress", len(request))
		// 入口带宽 接收响应
		kvs = append(kvs, "ingress", len(reply))

		log.Context(ctx).Log(level, kvs...)
		return nil
	}
}

func (logger *loggerHook) ProcessPipelineHook(next redisv9.ProcessPipelineHook) redisv9.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redisv9.Cmder) error {
		var (
			kvs   = make([]any, 0)
			level = log.LevelInfo
			err   error
		)

		if err = next(ctx, cmds); redisCmderError(err) != nil {
			level = log.LevelError
			kvs = append(kvs, "error", err.Error())
		}

		fn, file, line := funcFileLine("github.com/redis/go-redis")
		kvs = append(kvs, "msg",
			fmt.Sprintf("filepath: %s. lineno: %d. function: %s. numCmd: %d", file, line, fn, len(cmds)),
		)

		summary, cmdsString := rediscmd.CmdsString(cmds)
		kvs = append(kvs, "request", summary+cmdsString)

		log.Context(ctx).Log(level, kvs...)
		return err
	}
}

func funcFileLine(pkg string) (string, string, int) {
	const depth = 16
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	ff := runtime.CallersFrames(pcs[:n])

	var fn, file string
	var line int
	for {
		f, ok := ff.Next()
		if !ok {
			break
		}
		fn, file, line = f.Function, f.File, f.Line
		if !strings.Contains(fn, pkg) {
			break
		}
	}

	if ind := strings.LastIndexByte(fn, '/'); ind != -1 {
		fn = fn[ind+1:]
	}

	return fn, file, line
}

func redisCmderReply(cmd redisv9.Cmder) string {
	reply := cmd.String()
	splits := strings.SplitN(reply, ": ", 2)
	if len(splits) == 2 {
		return splits[1]
	}
	return reply
}

func redisCmderError(err error) error {
	switch err {
	case nil:
		return nil
	case redisv9.Nil:
		return nil
	default:
		return err
	}
}
