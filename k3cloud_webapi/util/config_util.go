package util

import (
	"bufio"
	"fmt"
	"my_project/go-demo/k3cloud_webapi/model"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// InitConfig 从配置文件初始化配置
func InitConfig(configPath, configNode string) (*model.ApiConfigS, error) {
	if configPath == "" {
		return nil, fmt.Errorf("Init config failed: Lack of config path!")
	}
	absPath := getAbsPath(configPath)
	if !fileExists(absPath) {
		return nil, fmt.Errorf("Init config failed: Config file[%s] not found!", absPath)
	}
	if configNode == "" {
		return nil, fmt.Errorf("Init config failed: Lack of config node!")
	}

	assignApiConfig := model.NewApiConfig()

	config, err := readConfig(absPath)
	if err != nil {
		return nil, err
	}

	nodeConfig, exists := config[configNode]
	if !exists {
		return assignApiConfig, nil
	}

	if serverURL := nodeConfig.get("x-kdapi-serverurl"); serverURL != "" {
		assignApiConfig.ServerURL = serverURL
	}
	assignApiConfig.Dcid = nodeConfig.get("x-kdapi-acctid")
	assignApiConfig.UserName = nodeConfig.get("x-kdapi-username")
	assignApiConfig.AppID = nodeConfig.get("x-kdapi-appid")
	assignApiConfig.AppSecret = nodeConfig.get("x-kdapi-appsec")

	assignApiConfig.Lcid = nodeConfig.getInt("x-kdapi-lcid", 2052)
	assignApiConfig.OrgNum = nodeConfig.getInt("x-kdapi-orgnum", 0)
	assignApiConfig.ConnectTimeout = nodeConfig.getInt("x-kdapi-connecttimeout", 120)
	assignApiConfig.RequestTimeout = nodeConfig.getInt("x-kdapi-requesttimeout", 120)
	assignApiConfig.Proxy = nodeConfig.get("x-kdapi-proxy")

	if secPwd := nodeConfig.get("x-kdapi-secpwd"); secPwd != "" {
		assignApiConfig.XorCode = secPwd
	}

	return assignApiConfig, nil
}

// InitConfigByParams 根据参数初始化配置
func InitConfigByParams(acctID, userName, appID, appSecret, serverURL string, lcid, orgNum, connectTimeout, requestTimeout int, proxy string) *model.ApiConfigS {
	assignApiConfig := model.NewApiConfig()

	if serverURL != "" {
		assignApiConfig.ServerURL = serverURL
	}
	assignApiConfig.Dcid = acctID
	assignApiConfig.UserName = userName
	assignApiConfig.AppID = appID
	assignApiConfig.AppSecret = appSecret

	if lcid > 0 {
		assignApiConfig.Lcid = lcid
	}
	if orgNum > 0 {
		assignApiConfig.OrgNum = orgNum
	}
	if connectTimeout > 0 {
		assignApiConfig.ConnectTimeout = connectTimeout
	}
	if requestTimeout > 0 {
		assignApiConfig.RequestTimeout = requestTimeout
	}
	assignApiConfig.Proxy = proxy

	return assignApiConfig
}

// getAbsPath 获取文件的绝对路径
func getAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return absPath
}

// fileExists 检查文件是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// readConfig 读取配置文件
func readConfig(filepath string) (map[string]configNode, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := make(map[string]configNode)
	var currentNode string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentNode = line[1 : len(line)-1]
			config[currentNode] = make(configNode)
		} else if currentNode != "" && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			value := strings.TrimSpace(parts[1])
			config[currentNode][key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

// configNode 定义配置节点
type configNode map[string]string

// get 获取配置节点中的值
func (cn configNode) get(key string) string {
	return cn[key]
}

// getInt 获取配置节点中的整数值
func (cn configNode) getInt(key string, defaultValue int) int {
	value, exists := cn[key]
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
