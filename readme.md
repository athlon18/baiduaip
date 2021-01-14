# baiduaip 实现人脸识别所使用的接口和通用文字识别（标准版）sdk

[![GoDoc](https://godoc.org/github.com/athlon18/baiduaip/baidu/aip?status.svg)](https://godoc.org/github.com/athlon18/baiduaip/baidu/aip)

```go
import (
    "github.com/athlon18/baiduaip/baidu/aip/client"
)

func main() {
    opt := newClientOptions()
    client.Init(opt)
    // ...
}

// newClientOptions 客户端选项
func newClientOptions() (opt *client.Option) {
    // ...
    return 
}

```


注: 运行`go test`前, 需要在baidu/aip/testdata目录下创建`key.json`和测试图片, 内容如下:

```json
{
    "AppID": "APPID",
    "APIKey": "APIKEY",
    "SecretKey": "SECRETKEY"
}
```

```shell
ls ./baidu/aip/testdata

cp.jpeg  EF2B8C2B8BCD43DA931B218759D59C22.jpeg  key.json  man4.jpg  tom.jpeg
```
