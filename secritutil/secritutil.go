package secritutil

import (
	"encoding/base64"
	"strings"

	"commonutils/crcutil"
)

// Encrypt 对给定的明文字符串加密
func Encrypt(dbInfo string) string {
	dbInfoHex := crcutil.TohexString([]byte(dbInfo))                                     // 明文字符串转16进制字符串
	encryptStr := base64.StdEncoding.EncodeToString([]byte(strings.Join(dbInfoHex, ""))) // 16进制字符串转base64加密字符串
	return encryptStr
}

// Decrypt 对给定的加密字符串解密
func Decrypt(encryptStr string) string {
	dbInfoByte, _ := base64.StdEncoding.DecodeString(encryptStr) // base64加密字符串转16进制字符串
	dbInfo := string(dbInfoByte)                                 // 16进制字符串转明文字符串
	return crcutil.ToNormalString(dbInfo)
}
