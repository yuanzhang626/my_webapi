package BD_Customer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"my_project/go-demo/k3cloud_webapi"
	"os"
	"strconv"
	"strings"
)

/*
获取所有产品信息
1.元数据接口
2.单据查询接口
3.查看接口
4.写文件
*/
var (
	fileName   = "output/客户列表.csv"
	formId     = "BD_Customer"
	listFields = []FieldInfo{
		{"FCUSTID", "实体主键"},
		{"FNumber", "客户编码"},
		{"FName", "客户名称"},
		{"FShortName", "简称"},
		{"FDocumentStatus", "单据状态"},
		{"FForbidStatus", "禁用状态"},
		{"FUseOrgId", "使用组织"},
		{"FGroup", "客户分组"},
		{"F_bog_CheckBox", "纯净水"},
		{"F_bog_CheckBox1", "矿泉水"},
		{"F_bog_CheckBox2", "山泉水"},
		{"F_bog_CheckBox3", "苏打水"},
		{"FCreatorId", "创建人"},
		{"FModifierId", "修改人"},
		{"FCreateDate", "创建日期"},
		{"FModifyDate", "修改日期"},
		{"FAPPROVEDATE", "审核日期"},
		{"FAPPROVERID", "审核人"},
	}
	fUseOrgId    = "1"      //使用组织，1 景田集团
	fCategoryID  = "198163" //供应商分组，
	filterString = []map[string]string{
		{
			"Left": "(", "FieldName": "FUseOrgId", "Compare": "=", "Value": fUseOrgId, "Right": ")", "Logic": "and",
		},
		//{
		//	"Left": "(", "FieldName": "FCategoryID", "Compare": "=", "Value": fCategoryID, "Right": ")", "Logic": "and",
		//},
	}
)

type FieldInfo struct {
	FieldKey  string
	FieldName string
}

func ListBDCustomer(sdk *k3cloud_webapi.K3CloudApiSdk) {
	//fmt.Println("hello")
	getAllListBDCustomer(sdk)
}

func listFieldsFieldKey(listFields []FieldInfo) []string {
	result := make([]string, 0, len(listFields))
	for _, field := range listFields {
		result = append(result, field.FieldKey)
	}
	return result
}

func listFieldsFieldName(listFields []FieldInfo) []string {
	result := make([]string, 0, len(listFields))
	for _, field := range listFields {
		result = append(result, field.FieldName)
	}
	return result
}

// 单据查询接口
func getAllListBDCustomer(sdk *k3cloud_webapi.K3CloudApiSdk) {
	fieldKeys := strings.Join(listFieldsFieldKey(listFields), ",")
	fieldNames := listFieldsFieldName(listFields)
	writeCsvFile(fileName, [][]string{fieldNames})

	startRow, limit := 0, 1000
	for {
		data, err := getListFromK3Cloud(sdk, formId, fieldKeys, filterString, startRow, limit)
		if err != nil {
			fmt.Println(err)
		}
		if len(data) == 0 {
			break
		}
		writeCsvFile(fileName, InterfaceSliceToStringSlice(data))
		startRow = startRow + limit
	}
}

func getListFromK3Cloud(sdk *k3cloud_webapi.K3CloudApiSdk, formId, fieldKeys string, filterString []map[string]string, startRow, limit int) (data [][]interface{}, err error) {
	para := map[string]interface{}{
		"FormId":       formId,
		"FieldKeys":    fieldKeys,
		"FilterString": filterString,
		"OrderString":  "",
		"TopRowCount":  0,
		"StartRow":     startRow,
		"Limit":        limit,
		"SubSystemId":  "",
	}

	response, err := sdk.ExecuteBillQuery(para)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		return nil, err
	}

	//fmt.Println(data)

	return data, nil
}

func writeCsvFile(filePath string, content [][]string) error {
	// 创建一个新的 CSV 文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640) //O_TRUNC	O_APPEND
	if err != nil {
		fmt.Println("创建文件时出错:", err)
		return err
	}
	defer file.Close()

	// 创建一个 CSV 写入器
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 一次性写入多行数据
	err = writer.WriteAll(content)
	if err != nil {
		fmt.Println("写入数据时出错:", err)
		return err
	}
	fmt.Printf("writeCsvFile %s %d rows success\n", filePath, len(content))
	return nil
}

func InterfaceSliceToStringSlice(in [][]interface{}) (out [][]string) {
	out = make([][]string, 0, len(in))

	for _, row := range in {
		outItem := make([]string, 0, len(row))
		for _, item := range row {
			switch v := item.(type) {
			case string:
				outItem = append(outItem, v)
			case int:
				outItem = append(outItem, strconv.Itoa(v))
			case bool:
				outItem = append(outItem, strconv.FormatBool(v))
			// 可以根据需要添加更多类型的转换
			default:
				// 处理其他类型，这里简单将其转换为字符串表示
				outItem = append(outItem, fmt.Sprintf("%v", v))
			}
		}
		out = append(out, outItem)
	}

	return
}
