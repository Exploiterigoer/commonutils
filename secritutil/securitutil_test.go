package secritutil

import (
	"encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Log(Encrypt("hello"))
}

func TestDecrypt(t *testing.T) {
	en := Encrypt("hello")
	t.Log(en)
	t.Log(base64.StdEncoding.DecodeString(en))
	t.Log(Decrypt(en))
}
