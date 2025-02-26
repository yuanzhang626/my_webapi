# k3cloud webapi
参考[官方sdk](https://openapi.open.kingdee.com/ApiSdkCenter)，使用golang来重写，后续将放到common库

主要工作
- 基础的http请求
- 封装成sdk

## 基础http请求



## sdk


## example
```golang
sdk := k3cloud_webapi.NewK3CloudApiSdk()
sdk.Init("./conf.ini", "config")
para := map[string]interface{}{
    "CreateOrgId": 0,
    "Number":      "",
    "Id":          "",
    "IsSortBySeq": "false",
}
response, err := sdk.View("BD_Material", para)
```
