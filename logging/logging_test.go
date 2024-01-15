package logging

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestZapLogging(t *testing.T) {
	log.NewHelper(log.WithContext(context.Background(), log.GetLogger())).Infof("test")
}
