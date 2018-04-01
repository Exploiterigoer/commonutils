// Package utf8ucsutil this package converts utf8 chars to ucs-2 or ucs-2 to utf8 char.
// 0000 0000-0000 007F | 0xxxxxxx                            | char with 1 byte
// 0000 0080-0000 07FF | 110xxxxx 10xxxxxx                   | char with 2 byte
// 0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx          | char with 3 byte
// 0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx | char with 4 byte
// char with 1 byte for testing:ASCII alphabet
// char with 2 byte for testing:Greek alphabet alphabet
// char with 2 byte for testing(\U0000-\UFFFF)：A variety of common Chinese chars and full width punctuation
// char with 2 byte for testing(\U00020000 ~ \U0002B81D)：𠀀、𪚥、𪚺、𠀁
// eg:
// UCS-2 code of "中国abcd" ---> 2D 4E FD 56 61 00 62 00 63 00 64 00
// UTF-8 code of "中国abcd" ---> E4 B8 AD E5 9B BD 61 62 63 64"
package utf8ucsutil

// ToUCS2 Converts utf8 string to ucs-2 code or ucs4 code
// if the a char with 4 byte,it will convert to ucs4 code,
// in others case,all wil convert to ucs2 code
// NOTE:
// the converting string must be a string with utf8 code
func ToUCS2(utf8Char string) []interface{} {
	utf8ByteArray := []byte(utf8Char)
	ucsArray := make([]interface{}, 0)
	length := len(utf8ByteArray)

	for i := 0; i < length; {
		if utf8ByteArray[i]>>7 == 0x00 {
			// utf-8 char with 1 byte
			ucsArray = append(ucsArray, uint16(utf8ByteArray[i])) // ucs-2 code consisted of 2 byte
			i++
		} else if utf8ByteArray[i]>>5 == 0x06 {
			// utf-8 char with 2 byte
			byte1 := uint16(utf8ByteArray[i]&0x1f) << 6 // last five bits of the first byte
			byte2 := uint16(utf8ByteArray[i+1] & 0x3f)  // last six bits of the second byte
			ucs2 := byte1 | byte2                       // ucs-2 code consisted of 2 byte
			ucsArray = append(ucsArray, uint16(ucs2))
			i += 2
		} else if utf8ByteArray[i]>>4 == 0x0e {
			// utf-8 char with 3 byte
			byte1 := uint16(utf8ByteArray[i]&0x0f) << 12  // last four bits of the first byte
			byte2 := uint16(utf8ByteArray[i+1]&0x3f) << 6 // last six bits of the second byte
			byte3 := uint16(utf8ByteArray[i+2] & 0x3f)    // last six bits of the third byte
			ucs2 := byte1 | byte2 | byte3                 // ucs-2 code consisted of 2 byte
			ucsArray = append(ucsArray, uint16(ucs2))
			i += 3
		} else {
			// utf-8 char with 4 byte
			byte1 := uint32(utf8ByteArray[i]&0x0f) << 18   // last three bits of the first byte
			byte2 := uint32(utf8ByteArray[i+1]&0x3f) << 12 // last six bits of the second byte
			byte3 := uint32(utf8ByteArray[i+2]&0x3f) << 6  // last six bits of the third byte
			byte4 := uint32(utf8ByteArray[i+3] & 0x3f)     // last six bits of the fourth byte
			ucs4 := byte1 | byte2 | byte3 | byte4          // ucs-2 code consisted of 4 byte
			ucsArray = append(ucsArray, uint32(ucs4))
			i += 4
		}
	}
	return ucsArray
}

// ToUTF8 Converts ucs code(ucs-2 or ucs-4) to utf8 char
// NOTE:
// every element of ucsArray is either the uint16 type or the uint32 type
// the uint16 type corresponds of ucs2,the uint32 type corresponds of ucs4
func ToUTF8(ucsArray []interface{}) string {
	char := make([]byte, 0)
	for _, val := range ucsArray {
		switch val.(type) {
		case uint16:
			v := val.(uint16)
			if v < 0x0080 {
				// utf-8 char with 1 byte
				char = append(char, byte(v))
			} else if v >= 0x0080 && v <= 0x07ff {
				// utf-8 char with 2 byte
				b1 := (v >> 6) | 0xc0
				b2 := (v & 0x003F) | 0x80
				char = append(char, byte(b1))
				char = append(char, byte(b2))
			} else if v >= 0x0800 && v <= 0xffff {
				// utf-8 char with 3 byte
				b1 := (v >> 12) | 0xe0
				b2 := ((v & 0x0fc0) >> 6) | 0x80
				b3 := (v & 0x003f) | 0x80

				char = append(char, byte(b1))
				char = append(char, byte(b2))
				char = append(char, byte(b3))
			}
		case uint32:
			// utf-8 char with 4 byte
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
