package k3cloud_webapi

import (
	"fmt"
	"my_project/go-demo/k3cloud_webapi/constant"
	"my_project/go-demo/k3cloud_webapi/core"
	"my_project/go-demo/k3cloud_webapi/util"
)

// K3CloudApiSdk 对应Python中的K3CloudApiSdk类
type K3CloudApiSdk struct {
	*core.WebApiClient
}

// NewK3CloudApiSdk 构造函数，对应Python中的__init__方法
func NewK3CloudApiSdk() *K3CloudApiSdk {
	sdk := &K3CloudApiSdk{
		core.NewWebApiClient(),
	}
	return sdk
}

// InitConfig 对应Python中的InitConfig方法
func (sdk *K3CloudApiSdk) InitConfig(acctID, userName, appID, appSecret, serverUrl string, lcid int, orgNum int, connectTimeout, requestTimeout int, proxy string) {
	sdk.Config = util.InitConfigByParams(acctID, userName, appID, appSecret, serverUrl, lcid, orgNum, connectTimeout, requestTimeout, proxy)
	sdkInitialize := sdk.IsValid()
	sdk.WebApiClient.Init(sdk.Config.ServerURL, 120, sdkInitialize)
}

// Init 对应Python中的Init方法
func (sdk *K3CloudApiSdk) Init(configPath, configNode string) {
	conf, err := util.InitConfig(configPath, configNode)
	if err != nil {
		panic(err)
	}
	sdk.Config = conf
	sdkInitialize := sdk.IsValid()

	sdk.WebApiClient.Init(sdk.Config.ServerURL, 120, sdkInitialize)
}

// IsValid 对应Python中的IsValid方法
func (sdk *K3CloudApiSdk) IsValid() bool {
	msg := ""
	if sdk.Config.Dcid == "" {
		msg += ",账套ID"
	}
	if sdk.Config.UserName == "" {
		msg += ",用户"
	}
	if sdk.Config.AppID == "" {
		msg += ",应用ID"
	}
	if sdk.Config.AppSecret == "" {
		msg += ",应用密钥"
	}
	if msg != "" {
		fmt.Printf("SDK初始化失败，缺少必填授权项：%s\n", msg[1:])
		return false
	}
	return true
}

// GetDataCenters 对应Python中的GetDataCenters方法
func (sdk *K3CloudApiSdk) GetDataCenters() (string, error) {
	return sdk.WebApiClient.Execute("Kingdee.BOS.ServiceFacade.ServicesStub.Account.AccountService.GetDataCenterList", nil, constant.SYNC)
}

// ExcuteOperation 对应Python中的ExcuteOperation方法
func (sdk *K3CloudApiSdk) ExcuteOperation(formid, opNumber string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid":   formid,
		"opnumber": opNumber,
		"data":     data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExcuteOperation", jsonData, constant.SYNC)
}

// Save 对应Python中的Save方法
func (sdk *K3CloudApiSdk) Save(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Save", jsonData, constant.SYNC)
}

// BatchSave 对应Python中的BatchSave方法
func (sdk *K3CloudApiSdk) BatchSave(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.BatchSave", jsonData, constant.SYNC)
}

// BatchSaveQuery 对应Python中的BatchSaveQuery方法
func (sdk *K3CloudApiSdk) BatchSaveQuery(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.BatchSave", jsonData, constant.QUERY)
}

// Audit 对应Python中的Audit方法
func (sdk *K3CloudApiSdk) Audit(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Audit", jsonData, constant.SYNC)
}

// Delete 对应Python中的Delete方法
func (sdk *K3CloudApiSdk) Delete(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Delete", jsonData, constant.SYNC)
}

// UnAudit 对应Python中的UnAudit方法
func (sdk *K3CloudApiSdk) UnAudit(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.UnAudit", jsonData, constant.SYNC)
}

// Submit 对应Python中的Submit方法
func (sdk *K3CloudApiSdk) Submit(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Submit", jsonData, constant.SYNC)
}

// View 对应Python中的View方法
func (sdk *K3CloudApiSdk) View(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.View", jsonData, constant.SYNC)
}

// ExecuteBillQuery 对应Python中的ExecuteBillQuery方法
func (sdk *K3CloudApiSdk) ExecuteBillQuery(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExecuteBillQuery", jsonData, constant.SYNC)
}

// BillQuery 对应Python中的BillQuery方法
func (sdk *K3CloudApiSdk) BillQuery(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.BillQuery", jsonData, constant.SYNC)
}

// Draft 对应Python中的Draft方法
func (sdk *K3CloudApiSdk) Draft(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Draft", jsonData, constant.SYNC)
}

// Allocate 对应Python中的Allocate方法
func (sdk *K3CloudApiSdk) Allocate(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Allocate", jsonData, constant.SYNC)
}

// FlexSave 对应Python中的FlexSave方法
func (sdk *K3CloudApiSdk) FlexSave(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.FlexSave", jsonData, constant.SYNC)
}

// SendMsg 对应Python中的SendMsg方法
func (sdk *K3CloudApiSdk) SendMsg(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.SendMsg", jsonData, constant.SYNC)
}

// Push 对应Python中的Push方法
func (sdk *K3CloudApiSdk) Push(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Push", jsonData, constant.SYNC)
}

// GroupSave 对应Python中的GroupSave方法
func (sdk *K3CloudApiSdk) GroupSave(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GroupSave", jsonData, constant.SYNC)
}

// Disassembly 对应Python中的Disassembly方法
func (sdk *K3CloudApiSdk) Disassembly(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Disassembly", jsonData, constant.SYNC)
}

// QueryBusinessInfo 对应Python中的QueryBusinessInfo方法
func (sdk *K3CloudApiSdk) QueryBusinessInfo(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.QueryBusinessInfo", jsonData, constant.SYNC)
}

// QueryGroupInfo 对应Python中的QueryGroupInfo方法
func (sdk *K3CloudApiSdk) QueryGroupInfo(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.QueryGroupInfo", jsonData, constant.SYNC)
}

// WorkflowAudit 对应Python中的WorkflowAudit方法
func (sdk *K3CloudApiSdk) WorkflowAudit(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.WorkflowAudit", jsonData, constant.SYNC)
}

// GroupDelete 对应Python中的GroupDelete方法
func (sdk *K3CloudApiSdk) GroupDelete(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GroupDelete", jsonData, constant.SYNC)
}

// SwitchOrg 对应Python中的SwitchOrg方法
func (sdk *K3CloudApiSdk) SwitchOrg(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.SwitchOrg", jsonData, constant.SYNC)
}

// cancelAllocate 对应Python中的cancelAllocate方法
func (sdk *K3CloudApiSdk) cancelAllocate(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAllocate", jsonData, constant.SYNC)
}

// cancelAssign 对应Python中的cancelAssign方法
func (sdk *K3CloudApiSdk) cancelAssign(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAssign", jsonData, constant.SYNC)
}

// getSysReportData 对应Python中的getSysReportData方法
func (sdk *K3CloudApiSdk) getSysReportData(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GetSysReportData", jsonData, constant.SYNC)
}

// attachmentUpload 对应Python中的attachmentUpload方法
func (sdk *K3CloudApiSdk) attachmentUpload(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentUpload", jsonData, constant.SYNC)
}

// attachmentDownLoad 对应Python中的attachmentDownLoad方法
func (sdk *K3CloudApiSdk) attachmentDownLoad(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentDownLoad", jsonData, constant.SYNC)
}

// CancelAllocate 对应Python中的CancelAllocate方法
func (sdk *K3CloudApiSdk) CancelAllocate(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAllocate", jsonData, constant.SYNC)
}

// CancelAssign 对应Python中的CancelAssign方法
func (sdk *K3CloudApiSdk) CancelAssign(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAssign", jsonData, constant.SYNC)
}

// GetSysReportData 对应Python中的GetSysReportData方法
func (sdk *K3CloudApiSdk) GetSysReportData(formid string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"formid": formid,
		"data":   data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GetSysReportData", jsonData, constant.SYNC)
}

// AttachmentUpload 对应Python中的AttachmentUpload方法
func (sdk *K3CloudApiSdk) AttachmentUpload(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentUpload", jsonData, constant.SYNC)
}

// AttachmentDownLoad 对应Python中的AttachmentDownLoad方法
func (sdk *K3CloudApiSdk) AttachmentDownLoad(data interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"data": data,
	}
	return sdk.WebApiClient.Execute("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentDownLoad", jsonData, constant.SYNC)
}

//func main() {
//	// 示例用法
//	sdk := NewK3CloudApiSdk("http://example.com", 120)
//	sdk.InitConfig("acct_id", "user_name", "app_id", "app_secret", "server_url", 2052, 0, 120, 120, "")
//	sdk.GetDataCenters()
//}
