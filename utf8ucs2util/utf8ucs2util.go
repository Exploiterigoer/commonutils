// utf8ucs2 这个包用于utf8字符与ucs-2字符编码的转换
// 0000 0000-0000 007F | 0xxxxxxx                            | 字符占1个字节
// 0000 0080-0000 07FF | 110xxxxx 10xxxxxx                   | 字符占2个字节
// 0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx          | 字符占3个字节
// 0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx | 字符占4个字节
// 1字节测试用的字符:ASCII字符
// 2字节测试用的字符:希腊字母、俄文字母等
// 3字节测试用的字符(\U0000-\UFFFF)：各种常见的汉字及全角标点符号
// 4字节测试用的字符(\U00020000 ~ \U0002B81D)：𠀀、𪚥、𪚺、𠀁
// "中国abcd"对应的UCS-2编码为"2D 4E FD 56 61 00 62 00 63 00 64 00"
// 对应的UTF-8编码为"E4 B8 AD E5 9B BD 61 62 63 64"

package utf8ucs2util

// ToUCS2 将utf8编码的字符串转换为ucs-2编码(2字节16位)
// 或ucs4编码(4字节32位)的字符串，ucs-2编码、ucs4编码
// 依次对应uint16类型、uint32类型的数据，因而返回interface结果
// 在使用前必须确保输入的参数时utf8编码
func ToUCS2(utf8Char string) []interface{} {
	utf8ByteArray := []byte(utf8Char)
	ucsArray := make([]interface{}, 0)
	length := len(utf8ByteArray)

	for i := 0; i < length; {
		if utf8ByteArray[i]>>7 == 0x00 {
			// 1个字节的utf-8编码的字符
			ucsArray = append(ucsArray, uint16(utf8ByteArray[i])) // 这里合成2个字节的ucs-2编码
			i++
		} else if utf8ByteArray[i]>>5 == 0x06 {
			// 2个字节的utf-8编码的字符
			byte1 := uint16(utf8ByteArray[i]&0x1f) << 6 // 这里取出第1个字节的后5位
			byte2 := uint16(utf8ByteArray[i+1] & 0x3f)  // 这里取出第2个字节的后6位
			ucs2 := byte1 | byte2                       // 这里合成2个字节的ucs-2编码
			ucsArray = append(ucsArray, uint16(ucs2))
			i += 2
		} else if utf8ByteArray[i]>>4 == 0x0e {
			// 3个字节的utf-8编码的字符
			byte1 := uint16(utf8ByteArray[i]&0x0f) << 12  // 这里取出第1个字节的后4位
			byte2 := uint16(utf8ByteArray[i+1]&0x3f) << 6 // 这里取出第2个字节的后6位
			byte3 := uint16(utf8ByteArray[i+2] & 0x3f)    // 这里取出第3个字节的后6位
			ucs2 := byte1 | byte2 | byte3                 // 这里合成2个字节的ucs-2编码
			ucsArray = append(ucsArray, uint16(ucs2))
			i += 3
		} else {
			// 4个字节的utf-8编码的字符
			byte1 := uint32(utf8ByteArray[i]&0x0f) << 18   // 这里取出第1个字节的后3位
			byte2 := uint32(utf8ByteArray[i+1]&0x3f) << 12 // 这里取出第2个字节的后6位
			byte3 := uint32(utf8ByteArray[i+2]&0x3f) << 6  // 这里取出第3个字节的后6位
			byte4 := uint32(utf8ByteArray[i+3] & 0x3f)     // 这里取出第4个字节的后6位
			ucs4 := byte1 | byte2 | byte3 | byte4          // 这里合成4个字节的ucs-4编码
			ucsArray = append(ucsArray, uint32(ucs4))
			i += 4
		}
	}
	return ucsArray
}

// ToUTF8 将ucs(ucs-2或ucs-4)编码的字符转换为utf-8编码的字符串
// ucsArray 是字符对应的ucs编码数组，每个元素要么数uint16，要么是
// uint32，uint16类型的元素对应ucs-2字符的编码，uint32类型则的对
// 应ucs-4字符的编码
func ToUTF8(ucsArray []interface{}) string {
	char := make([]byte, 0)
	for _, val := range ucsArray {
		switch val.(type) {
		case uint16:
			v := val.(uint16)
			if v < 0x0080 {
				// 占1个字节的utf8字符
				char = append(char, byte(v))
			} else if v >= 0x0080 && v <= 0x07ff {
				// 占2个字节的utf8字符
				b1 := (v >> 6) | 0xc0
				b2 := (v & 0x003F) | 0x80
				char = append(char, byte(b1))
				char = append(char, byte(b2))
			} else if v >= 0x0800 && v <= 0xffff {
				// 占3个字节的utf8字符
				b1 := (v >> 12) | 0xe0
				b2 := ((v & 0x0fc0) >> 6) | 0x80
				b3 := (v & 0x003f) | 0x80

				char = append(char, byte(b1))
				char = append(char, byte(b2))
				char = append(char, byte(b3))
			}
		case uint32:
			// 占4个字节的utf8字符
			v := val.(uint32)
			b1 := (v >> 18) | 0xf0
			b2 := (v >> 12) | 0x80
			b3 := (v >> 6) | 0x80
			b4 := (v & 0x003f) | 0x80
			char = append(char, byte(b1))
			char = append(char, byte(b2))
			char = append(char, byte(b3))
			char = append(char, byte(b4))
		}
	}
	return string(char)
}
