package secritutil

import (
	"encoding/base64"
	"strings"

	"github.com/Exploiterigoer/commonutils/byteutil"
)

// Encrypt Encodes the given string
func Encrypt(dbInfo string) string {
	// hexing of the given string
	dbInfoHex := byteutil.TohexString([]byte(dbInfo))
	// encoding the  hex  string
	encryptStr := base64.StdEncoding.EncodeToString([]byte(strings.Join(dbInfoHex, "")))
	return encryptStr
}

// Decrypt Decodes the given string
func Decrypt(encryptStr string) string {
	// decoding the given encoding string to non-crypto hexing string
	dbInfoByte, _ := base64.StdEncoding.DecodeString(encryptStr)
	// hexing string to normal string
	dbInfo := string(dbInfoByte)
	return byteutil.ToNormalString(dbInfo)
}
