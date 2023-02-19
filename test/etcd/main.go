package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 链接ETCD
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	leaseID, err := lease.Grant(context.TODO(), 10) // 生成一个租约id   10秒过期
	if err != nil {
		panic(err)
	}

	// 我们在这里存储一个值
	put, err := kv.Put(context.TODO(), "key1", "val1", clientv3.WithLease(leaseID.ID))
	if err != nil {
		panic(err)
	}
	put, err = kv.Put(context.TODO(), "key2", "val1", clientv3.WithLease(leaseID.ID))
	if err != nil {
		panic(err)
	}
	put, err = kv.Put(context.TODO(), "key3", "val1", clientv3.WithLease(leaseID.ID))
	if err != nil {
		panic(err)
	}

	fmt.Println(put)

	// 自动定时续租
	_, err = lease.KeepAlive(context.TODO(), leaseID.ID)
	if err != nil {
		panic(err)
	}

	// 循环查询
	for {
		time.Sleep(time.Second)
		get, err := kv.Get(context.TODO(), "key", clientv3.WithPrefix())
		if err != nil {
			panic(err)
		}
		fmt.Println(get)
	}
}
