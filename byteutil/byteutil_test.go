package byteutil

import (
	"testing"
)

func TestPadZero(t *testing.T) {
	t.Log(PadZero(10))
}

func TestTohexString(t *testing.T) {
	t.Log(TohexString([]byte("hello")))
}

func TestByteToInt(t *testing.T) {
	t.Log(ByteToInt([]byte{0x01, 0x02, 0x03, 0x04}, 1, 16))
}

func TestIntToByte(t *testing.T) {
	t.Log(IntToByte(258, 1, 4))
}

func TestByteToUCS2(t *testing.T) {
	t.Log(ByteToUCS2([]byte("hello")))
}

func TestUCS2ToByte(t *testing.T) {
	b := []byte("hello")
	t.Log(b)

	uc := ByteToUCS2(b)
	b1 := make([]byte, 0)
	for i := range uc {
		u := UCS2ToByte(uc[i].(uint16), 1)
		b1 = append(b1, u...)
	}
	t.Log(string(b1))
}
