package util

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"my_project/go-demo/k3cloud_webapi/model"
	"strconv"
	"time"
)

// DecodeAppSecret 解码AppSecret
func DecodeAppSecret(appSecret string) string {
	if len(appSecret) != 32 {
		return ""
	}
	base64Decode, _ := base64.StdEncoding.DecodeString(appSecret)
	base64Xor := xorCode(base64Decode)
	return base64.StdEncoding.EncodeToString(base64Xor)
}

// extendByteArray 扩展字节数组
func extendByteArray(origin string, encoding string, extendType int) []byte {
	if extendType == 0 {
		return []byte(rot(origin))
	}
	geneStr := ""
	for i := 0; i < 4; i++ {
		geneStr += origin[i*9 : i*9+8]
	}
	return []byte(rot(geneStr))
}

// xorCode 异或编码
func xorCode(byteArray []byte) []byte {
	pwdArray := extendByteArray(generateCode(), "utf-8", 1)
	outArray := make([]byte, len(byteArray))
	for i := 0; i < len(byteArray); i++ {
		outArray[i] = byteArray[i] ^ pwdArray[i]
	}
	return outArray
}

// encodeChar 编码字符
func encodeChar(ch byte) byte {
	x := byte(97)
	if ch >= 'A' && ch <= 'Z' {
		x = 65
	} else if ch < 'a' || ch > 'z' {
		return ch
	}
	return byte((int(ch)-int(x)+13)%26 + int(x))
}

// rot 字符旋转编码
func rot(s string) string {
	var result bytes.Buffer
	for _, c := range s {
		result.WriteByte(encodeChar(byte(c)))
	}
	return result.String()
}

// generateCode 生成代码
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	randNum := strconv.Itoa(rand.Intn(9000) + 1000)
	retCode := ""
	if model.ApiConfig.XorCode == "" {
		retCode += "0054s397" + string(randNum[0])
		retCode += "p6234378" + string(randNum[1])
		retCode += "o09pn7q3" + string(randNum[2])
		retCode += "r5qropr7" + string(randNum[3])
	} else {
		for i := 0; i < 4; i++ {
			retCode += model.ApiConfig.XorCode[i*8:(i+1)*8] + string(randNum[i])
		}
	}
	return retCode
}
