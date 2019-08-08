package discovery

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

var Endpoints = []string{"http://127.0.0.1:2379"}
var TTL int64 = 5

func ConnectEtcd() (etcdClient *clientv3.Client, err error) {
	config := clientv3.Config{
		Endpoints:         Endpoints,
		DialTimeout:       time.Second * 30,
		DialKeepAliveTime: time.Second * 30,
	}
	etcdClient, err = clientv3.New(config)
	return
}
