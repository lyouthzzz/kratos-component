package etcd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_Build(t *testing.T) {
	c := DefaultConfig().WithEndpoints("localhost:2379").WithTimeout("2s").WithBasicAuth("root", "root")
	require.Equal(t, []string{"localhost:2379"}, c.Endpoints)
	require.Equal(t, "2s", c.Timeout)
	require.Equal(t, "root", c.Username)
	require.Equal(t, "root", c.Password)
}
