// 对应 Python 中的 HeaderParam 类
package constant

var HeaderParam *HeaderParamS

func init() {
	HeaderParam = &HeaderParamS{
		X_Api_ClientID:      "X-Api-ClientID",
		X_Api_Auth_Version:  "X-Api-Auth-Version",
		X_Api_Timestamp:     "x-api-timestamp",
		X_Api_Nonce:         "x-api-nonce",
		X_Api_SignHeaders:   "x-api-signheaders",
		X_Api_Signature:     "X-Api-Signature",
		X_KD_AppKey:         "X-Kd-Appkey",
		X_KD_AppData:        "X-Kd-Appdata",
		X_KD_Signature:      "X-Kd-Signature",
		KDService_SessionId: "kdservice-sessionid",
		Cookie_Set:          "Set-Cookie",
	}
}

// HeaderParam 定义了一些常量，用于表示HTTP请求头的键名
type HeaderParamS struct {
	X_Api_ClientID      string
	X_Api_Auth_Version  string
	X_Api_Timestamp     string
	X_Api_Nonce         string
	X_Api_SignHeaders   string
	X_Api_Signature     string
	X_KD_AppKey         string
	X_KD_AppData        string
	X_KD_Signature      string
	KDService_SessionId string
	Cookie_Set          string
}
