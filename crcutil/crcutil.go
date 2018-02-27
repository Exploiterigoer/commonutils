// crc16校验码生成工具
package crcutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// TohexString 字节转换为16进制字符串
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

// ToNormalString 16进制字符转为常规字符
func ToNormalString(hexStr string) string {
	hexs := strings.Split(hexStr, "")
	hexb := make([]byte, 0)
	var hb int64
	tmp := ""
	for k, v := range hexs {
		if k%2 == 0 && len(tmp) == 2 {
			hb, _ = strconv.ParseInt(tmp, 16, 32)
			hexb = append(hexb, byte(hb))
			tmp = ""
		}
		tmp += v
	}
	// 最后一次循环，取2余数不为0的
	hb, _ = strconv.ParseInt(tmp, 16, 32)
	hexb = append(hexb, byte(hb))
	return string(hexb)
}

//将低序字节存储在起始地址，俗称小端，低8位
//将高序字节存储在起始地址，俗称大端，高8位
//通用modbus CRC校验算法
func modbusCRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//通用modbus取寄存器的低8位参与异或运算
		crc = ((crc << 8) >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//将得到的校验码按照先高字节后低字节的顺序存放（小端）
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, int16(crc))
	hex := TohexString(bytesBuffer.Bytes())
	return hex[0] + hex[1]
}

//HJ212 CRC校验算法
func hjt212CRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//国标取寄存器的高8位参与异或运算
		crc = (crc >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//将得到的校验码按照先高字节后低字节的顺序存放（小端）
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, int16(crc))
	hex := TohexString(bytesBuffer.Bytes())
	return hex[1] + hex[0]
}

//两种算法的计算结果
func CRC16(dataString string) (string, string) {
	return hjt212CRC(dataString), modbusCRC(dataString)
}
