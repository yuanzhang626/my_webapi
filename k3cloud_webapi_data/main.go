package main

import (
	"my_project/go-demo/k3cloud_webapi"
	"my_project/go-demo/k3cloud_webapi_data/BD_Customer"
	"my_project/go-demo/k3cloud_webapi_data/BD_Material"
	"my_project/go-demo/k3cloud_webapi_data/BD_STOCK"
	"my_project/go-demo/k3cloud_webapi_data/BD_Supplier"
)

func main() {
	//fmt.Println("hello world")
	sdk := k3cloud_webapi.NewK3CloudApiSdk()
	sdk.Init("./conf.ini", "config")
	//kwargs := map[string]interface{}{
	//	"Number": "1.01.01100001",
	//}
	//MaterialView(sdk, kwargs)
	if true {
		BD_Material.ListBDMaterial(sdk)
		BD_Supplier.ListBDSupplier(sdk)
		BD_Customer.ListBDCustomer(sdk)
	}

	BD_STOCK.ListBDSTOCK(sdk)

}
