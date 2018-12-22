# Golang Notebook

## 生成coverage html file

```bash
#!/bin/bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## net

### 获取到随机端口的TCP地址

参考示例：

go-nsq/consumer_test.go#consumerTest()方法:

```go
laddr := "127.0.0.1"
// so that the test can simulate binding consumer to specified address
config.LocalAddr, _ = net.ResolveTCPAddr("tcp", laddr+":0")
```



