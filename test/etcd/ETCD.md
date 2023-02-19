## ETCD  

好久都没有用了忘记了 好多 

我们这里几个内部服务都通过ETCD来做服务发现与注册 我们会利用ETCD的 租约特性

`go get go.etcd.io/etcd/client/v3`


// 服务注册   服务名称-随机id
``` 
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
```

// 心跳
```  
    // 自动定时续租
	_, err = lease.KeepAlive(context.TODO(), leaseID.ID)
	if err != nil {
		panic(err)
	}
```

// 服务发现 通过前缀查询
``` 
		get, err := kv.Get(context.TODO(), "key", clientv3.WithPrefix())
		if err != nil {
			panic(err)
		}
		fmt.Println(get)
```