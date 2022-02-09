package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

//go:generate protoc -I. --go_out=paths=source_relative:. etcd_config.proto

func DefaultConfig() *Config {
	return &Config{Timeout: "3s"}
}

func (x *Config) WithEndpoints(endpoints ...string) *Config {
	x.Endpoints = endpoints
	return x
}

func (x *Config) WithTimeout(timeout string) *Config {
	x.Timeout = timeout
	return x
}

func (x *Config) WithUsername(username string) *Config {
	x.Username = username
	return x
}

func (x *Config) WithPassword(password string) *Config {
	x.Password = password
	return x
}

func (x *Config) WithBasicAuth(username, password string) *Config {
	x.Username = username
	x.Password = password
	return x
}

func (x *Config) Build() (*clientv3.Client, error) {
	timeout, err := time.ParseDuration(x.Timeout)
	if err != nil {
		return nil, err
	}
	ec := clientv3.Config{Endpoints: x.Endpoints, Username: x.Username, Password: x.Password, DialTimeout: timeout}
	return clientv3.New(ec)
}

func (x *Config) BuildMust() *clientv3.Client {
	client, err := x.Build()
	if err != nil {
		panic(err)
	}
	return client
}
