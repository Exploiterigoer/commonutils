// Package crcutil a tool of CRC16 checkout algorithm.
package crcutil

import (
	"bytes"
	"encoding/binary"

	"github.com/Exploiterigoer/commonutils/byteutil"
)

// modbusCRC modbus CRC16 of checkout algorithm.
func modbusCRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		// gets the low 8 bits for calculating.
		crc = crc ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//formats the check code with littleEndian.
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, int16(crc))
	hex := byteutil.TohexString(bytesBuffer.Bytes())
	return hex[0] + hex[1]
}

//HJ212 CRC16 of checkout algorithm.
func hjt212CRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		// gets the heigh 8 bits for calculating.
		crc = (crc >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//formats the check code with littleEndian.
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, int16(crc))
	hex := byteutil.TohexString(bytesBuffer.Bytes())
	return hex[1] + hex[0]
}

//CRC16 the result of the two checkout algrothim method.
func CRC16(dataString string) (string, string) {
	return hjt212CRC(dataString), modbusCRC(dataString)
}
