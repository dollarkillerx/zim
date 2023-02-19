# cmd

``` 
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
    
    只需要改动 helloworld/helloworld.proto 为制定文件即可
```


// https://www.lixueduan.com/posts/grpc/06-auth/
