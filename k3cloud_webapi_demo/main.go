package main

import (
	"fmt"
	"my_project/go-demo/k3cloud_webapi"
)

func main() {
	//fmt.Println("hello world")
	sdk := k3cloud_webapi.NewK3CloudApiSdk()
	sdk.Init("./conf.ini", "config")
	kwargs := map[string]interface{}{
		"Number": "1.01.01100001",
	}
	MaterialView(sdk, kwargs)

}

func MaterialView(sdk *k3cloud_webapi.K3CloudApiSdk, kwargs map[string]interface{}) string {
	fmt.Println("[yuan]物料查看接口 enter")
	// 定义默认参数
	para := map[string]interface{}{
		"CreateOrgId": 0,
		"Number":      "",
		"Id":          "",
		"IsSortBySeq": "false",
	}
	// 如果传入了额外参数，更新默认参数
	if len(kwargs) > 0 {
		for key, value := range kwargs {
			para[key] = value
		}
	}
	fmt.Println("[yuan]物料查看接口 api_sdk.View before")
	// 调用 ApiSdk 的 View 方法获取响应
	response, err := sdk.View("BD_Material", para)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[yuan]物料查看接口 api_sdk.View after")
	fmt.Printf("物料查看接口：%s\n", response)
	// 调用 CheckResponse 函数检查响应
	return response
}
