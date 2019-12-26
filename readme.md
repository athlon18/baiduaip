# baiduaip 实现人脸识别所使用的接口sdk

[![GoDoc](https://godoc.org/github.com/antlinker/baiduaip/baidu/aip?status.svg)](https://godoc.org/github.com/antlinker/baiduaip/baidu/aip)

```go
import (
    "github.com/antlinker/baiduaip/baidu/aip/client"
)

func main() {
    // InitBaiduFaceOptions 初始化百度人脸服务的客户端
func InitBaiduFaceOptions(opt *client.Option) {
    client.Init(opt)
    // ...
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
