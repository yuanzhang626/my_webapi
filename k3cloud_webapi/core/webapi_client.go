package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"my_project/go-demo/k3cloud_webapi/constant"
	"my_project/go-demo/k3cloud_webapi/model"
	"my_project/go-demo/k3cloud_webapi/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ValidResult 验证结果
func ValidResult(responseContent string) (string, error) {
	if strings.HasPrefix(responseContent, "response_error:") {
		resError := strings.TrimSpace(strings.TrimPrefix(responseContent, "response_error:"))
		if resError != "" {
			return "", fmt.Errorf(resError)
		} else {
			return "", fmt.Errorf("Empty exception message")
		}
	}
	return responseContent, nil
}

// WebApiClient Web API客户端
type WebApiClient struct {
	Initialize     bool
	Config         *model.ApiConfigS
	Identify       *model.IdentifyS
	cookiesStore   *model.CookieStore
	connectTimeout time.Duration
	requestTimeout time.Duration
	proxy          string
}

// NewWebApiClient 创建新的WebApiClient实例
func NewWebApiClient() *WebApiClient {
	return &WebApiClient{
		Initialize:     false,
		Config:         nil, //model.NewApiConfig(),
		Identify:       nil,
		cookiesStore:   model.NewCookieStore("", nil),
		connectTimeout: 120 * time.Second,
		requestTimeout: 120 * time.Second,
		proxy:          "",
	}
}

// Init 初始化
func (c *WebApiClient) Init(serverURL string, timeout int, sdkInitialize bool) *WebApiClient {
	c.Initialize = sdkInitialize
	if c.Config == nil {
		panic("Configuration file unload, you need initial Config firstly!")
	}
	if serverURL == "" {
		serverURL = c.Config.ServerURL
	}
	// 这里需要根据config结构调整
	c.Identify = model.NewIdentify(serverURL, c.Config.Dcid, c.Config.UserName, c.Config.AppID, c.Config.AppSecret, c.Config.OrgNum, c.Config.Lcid, "")

	if timeout > 0 {
		c.requestTimeout = time.Duration(timeout) * time.Second
	}
	// 这里需要根据config结构调整
	if c.Config.Proxy != "" {
		c.proxy = c.Config.Proxy
	}
	return c
}

// BuildHeader 构建请求头
func (c *WebApiClient) BuildHeader(serviceURL string) map[string]string {
	pathURL := serviceURL
	if strings.HasPrefix(serviceURL, "http") {
		if len(serviceURL) > 10 {
			serviceURL = serviceURL[10:]
			pIndex := strings.Index(serviceURL, "/")
			if pIndex > -1 {
				pathURL = serviceURL[pIndex:]
			}
		}

	}
	pathURL = url.QueryEscape(pathURL)
	pathURL = strings.ReplaceAll(pathURL, "/", "%2F")
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	clientID := ""
	clientSec := ""
	arr := strings.Split(c.Identify.AppId, "_")
	if len(arr) == 2 {
		clientID = arr[0]
		clientSec = util.DecodeAppSecret(arr[1])
	}

	apiSign := "POST\n" + pathURL + "\n\nx-api-nonce:" + nonce + "\nx-api-timestamp:" + timeStamp + "\n"
	appData := fmt.Sprintf("%s,%s,%d,%d", c.Identify.DCID, c.Identify.UserName, c.Identify.LCID, c.Identify.OrgNum)
	//c.Identify.DCID + "," + c.Identify.UserName + "," + c.Identify.LCID + "," + c.Identify.OrgNum

	dicHeader := map[string]string{
		constant.HeaderParam.X_Api_ClientID:     clientID,
		constant.HeaderParam.X_Api_Auth_Version: "2.0",
		constant.HeaderParam.X_Api_Timestamp:    timeStamp,
		constant.HeaderParam.X_Api_Nonce:        nonce,
		constant.HeaderParam.X_Api_SignHeaders:  "x-api-timestamp,x-api-nonce",
		constant.HeaderParam.X_Api_Signature:    hmacSHA256(apiSign, clientSec),
		constant.HeaderParam.X_KD_AppKey:        c.Identify.AppId,
		constant.HeaderParam.X_KD_AppData:       base64.StdEncoding.EncodeToString([]byte(appData)),
		constant.HeaderParam.X_KD_Signature:     hmacSHA256(c.Identify.AppId+appData, c.Identify.AppSecret),
	}

	if c.cookiesStore.SID != "" {
		dicHeader[constant.HeaderParam.KDService_SessionId] = c.cookiesStore.SID
	}
	if len(c.cookiesStore.Cookies) > 0 {
		cookieStr := "Theme=standard"
		for _, v := range c.cookiesStore.Cookies {
			cookieStr += ";" + v.Name + "=" + v.Value
		}
		dicHeader["Cookie"] = cookieStr
	}

	dicHeader["Accept-Charset"] = "utf-8"
	dicHeader["User-Agent"] = "Kingdee/Go WebApi SDK 7.3 (compatible; MSIE 6.0; Windows NT 5.1;SV1)"
	dicHeader["Content-Type"] = "application/json"

	return dicHeader
}

// PostJson 发送POST请求
func (c *WebApiClient) PostJson(serviceName string, jsonData map[string]interface{}, invokeType constant.InvokeMethod) (string, error) {
	fmt.Println("[yuan] PostJson enter", serviceName, jsonData, invokeType)
	if jsonData == nil {
		jsonData = make(map[string]interface{})
	}
	reqURL := c.Identify.ServerUrl
	if strings.HasSuffix(reqURL, "/") {
		reqURL += serviceName + ".common.kdsvc"
	} else {
		reqURL += "/" + serviceName + ".common.kdsvc"
	}

	proxies := ""
	if c.proxy != "" {
		//todo 代理参数处理
		proxies = c.proxy
	}

	if invokeType == constant.QUERY {
		// 这里需要实现QueryMode常量
		jsonData[constant.BeginMethod_Header] = constant.BeginMethod_Method
		jsonData[constant.QueryMethod_Header] = constant.QueryMethod_Method
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: c.requestTimeout,
	}
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	for k, v := range c.BuildHeader(reqURL) {
		req.Header.Set(k, v)
	}

	if proxies != "" {
		proxyURL, err := url.Parse(proxies)
		if err != nil {
			return "", err
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	fmt.Println("[yuan] requests.post before url", reqURL)
	fmt.Println("[yuan] requests.post before headers", c.BuildHeader(reqURL))
	fmt.Println("[yuan] requests.post before data", string(jsonBytes))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("[yuan] requests.post after", resp.Status)

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusPartialContent {
		c.FillCookieAndHeader(resp.Cookies(), resp.Header)
		body := make([]byte, resp.ContentLength)
		_, err := resp.Body.Read(body)
		if err != nil {
			return "", err
		}
		return ValidResult(string(body))
	} else {
		return "", fmt.Errorf(resp.Status)
	}
}

// FillCookieAndHeader 填充cookie和header
func (c *WebApiClient) FillCookieAndHeader(cookies []*http.Cookie, headers http.Header) {
	for _, cookie := range cookies {
		if cookie.Name == constant.HeaderParam.KDService_SessionId {
			c.cookiesStore.SID = cookie.Value
		}
	}
	if _, ok := headers[constant.HeaderParam.Cookie_Set]; ok {
		c.cookiesStore.Cookies = make(map[string]*model.Cookie)
		cookieSet := headers[constant.HeaderParam.Cookie_Set][0]
		cookieParts := strings.Split(cookieSet, ",")
		for _, part := range cookieParts {
			fmt.Printf(part)
			ck := model.Parse(part)
			if ck != nil {
				c.cookiesStore.Cookies[ck.Name] = ck
			}
		}
	}
}

// Execute 执行请求
func (c *WebApiClient) Execute(serviceName string, jsonData map[string]interface{}, invokeType constant.InvokeMethod) (string, error) {
	fmt.Println("[yuan] Execute enter", serviceName, jsonData, invokeType)
	if !c.Initialize {
		return "", fmt.Errorf("拒绝请求，请先正确初始化!")
	}
	if c.Config == nil {
		return "", fmt.Errorf("请先初始化SDK配置信息!")
	}
	if invokeType == constant.SYNC {
		return c.PostJson(serviceName, jsonData, constant.SYNC)
	} else if invokeType == constant.QUERY {
		return c.ExecuteByQuery(serviceName, jsonData)
	} else {
		return "", fmt.Errorf("Not support for InvokeMode:" + strconv.Itoa(int(invokeType)))
	}
}

// ExecuteByQuery 通过查询执行
func (c *WebApiClient) ExecuteByQuery(serviceName string, jsonData map[string]interface{}) (string, error) {
	responseContent, err := c.PostJson(serviceName, jsonData, constant.QUERY)
	if err != nil {
		return "", err
	}
	var jsonResult map[string]interface{}
	err = json.Unmarshal([]byte(responseContent), &jsonResult)
	if err != nil {
		return "", err
	}

	if jsonResult["Status"] == constant.Complete {
		result, _ := json.Marshal(jsonResult["Result"])
		return string(result), nil
	} else {
		return c.QueryTaskResult(serviceName, map[string]interface{}{"TaskId": jsonResult["TaskId"], "Cancelled": false}, 5)
	}
}

// QueryTaskResult 查询任务结果
func (c *WebApiClient) QueryTaskResult(serviceName string, param map[string]interface{}, retryCount int) (string, error) {
	time.Sleep(1 * time.Second)
	queryService := serviceName[:strings.LastIndex(serviceName, ".")] + "." + constant.QueryMethod_Method
	queryInfo, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	responseContent, err := c.PostJson(queryService, map[string]interface{}{"queryInfo": string(queryInfo)}, constant.SYNC)
	if err != nil {
		if retryCount > 0 {
			return c.QueryTaskResult(serviceName, param, retryCount-1)
		} else {
			return "", err
		}
	}
	var jsonResult map[string]interface{}
	err = json.Unmarshal([]byte(responseContent), &jsonResult)
	if err != nil {
		return "", err
	}

	if jsonResult["Status"] == constant.Complete {
		result, _ := json.Marshal(jsonResult["Result"])
		return string(result), nil
	} else {
		return c.QueryTaskResult(serviceName, param, 5)
	}
}

func hmacSHA256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	signHex := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(signHex))
}
