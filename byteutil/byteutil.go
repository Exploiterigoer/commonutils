// 整数字节互转，时间提取，ID序列号生成工具
package byteutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// PadZero 补零
// count 填充的0的个数
func PadZero(count int) []byte {
	zero := make([]byte, 0)
	for i := 0; i < count; i++ {
		zero = append(zero, 0x00)
	}
	return zero
}

// TohexString 将字节数组转为16进制显示
// msg 被转换的字节数组
func TohexString(msg []byte) []string {
	hexString := make([]string, len(msg))
	for k, v := range msg {
		tmp := fmt.Sprintf("%X", v)
		if len(tmp) < 2 {
			tmp = "0" + tmp
		}
		hexString[k] = tmp
	}
	return hexString
}

// ByteToInt1 按照指定的大小端和正整数类型转换为整数
// ord 大小端对齐的顺序,1表示大端,0表示小端
// intType 转换的参照类型,为16,32,64任一值
// 16对应2字节的整数
// 32对应4字节的整数
// 64对应8字节的整数
func ByteToInt(b []byte, ord, intType int) uint64 {
	if intType < 16 {
		intType = 16
	}
	var i uint64
	switch intType {
	case 16:
		var tmp uint16
		if ord == 1 {
			binary.Read(bytes.NewBuffer(b), binary.BigEndian, &tmp)
		} else {
			binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &tmp)
		}
		i = uint64(tmp)
	case 32:
		var tmp uint32
		if ord == 1 {
			binary.Read(bytes.NewBuffer(b), binary.BigEndian, &tmp)
		} else {
			binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &tmp)
		}
		i = uint64(tmp)
	case 64:
		var tmp uint64
		if ord == 1 {
			binary.Read(bytes.NewBuffer(b), binary.BigEndian, &tmp)
		} else {
			binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &tmp)
		}
		i = tmp
	}
	return i
}

// IntToByte 将指定的整数转为指定大小端的字节数组
// i 被转换的整数
// ord 大小端对齐的顺序,1表示大端,0表示小端
// bitSize 字节的长度,1,4,8中的任一值
func IntToByte(i, ord, bitSize int) []byte {
	buffer := bytes.NewBuffer([]byte{})
	switch bitSize {
	case 1: // 1字节大小的整数
		tmp := uint8(i)
		if ord == 1 {
			binary.Write(buffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(buffer, binary.LittleEndian, &tmp)
		}
	case 4: // 4字节大小的整数
		tmp := uint32(i)
		if ord == 1 {
			binary.Write(buffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(buffer, binary.LittleEndian, &tmp)
		}
	case 8: // 8字节大小的整数
		tmp := uint64(i)
		if ord == 1 {
			binary.Write(buffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(buffer, binary.LittleEndian, &tmp)
		}
	}
	return buffer.Bytes()
}

// ByteToUCS2 将字节数组转为UCS2数组
// b 被转换的字节数组
// UCS2编码的字符规定2个字节
func ByteToUCS2(b []byte) []interface{} {
	ucs2 := make([]interface{}, 0)
	for k, _ := range b {
		if (k+2)%2 == 0 {
			x := ByteToInt(b[k:k+2], 1, 16) // 网络传输使用大端顺序,接收机器解析使用小端顺序
			ucs2 = append(ucs2, uint16(x))
		}
	}
	return ucs2
}

// UCS2ToByte 将UCS2的字符编码装维字节数组
// u UCS2的字符编码
// ord 大小端标记,1表示大端,0表示小端
func UCS2ToByte(u uint16, ord int) []byte {
	buffer := bytes.NewBuffer([]byte{})
	if ord == 1 {
		binary.Write(buffer, binary.BigEndian, u)
	} else {
		binary.Write(buffer, binary.LittleEndian, u)
	}
	return buffer.Bytes()
}
