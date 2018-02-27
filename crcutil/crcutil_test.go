package crcutil

import (
	"testing"
)

func TestCRC16(t *testing.T) {
	t.Log(CRC16("hello"))
}
