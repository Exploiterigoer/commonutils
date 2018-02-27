package utf8ucs2util

import (
	"testing"
)

func TestToUCS2(t *testing.T) {
	str := "你hello好"
	t.Log(ToUCS2(str))
}

func TestToUTF8(t *testing.T) {
	str := "你hello好奤"
	uc := ToUCS2(str)
	t.Log(uc)
	uf := ToUTF8(uc)
	t.Log(uf)
}
