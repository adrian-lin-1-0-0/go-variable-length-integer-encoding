package encode

import (
	"encoding/binary"
)

const (
	oneByteMax   uint64 = 0x3f
	twoByteMax   uint64 = 0x3fff
	fourByteMax  uint64 = 0x3fffffff
	eightByteMax uint64 = 0x3fffffffffffffff
)

type Error string

func (err Error) Error() string { return string(err) }

const (
	unknownErr = Error("unknown error")
)

func ToVarLenInt(num uint64) ([]byte, error) {
	if num <= oneByteMax {
		return encodeOneByte(uint8(num)), nil
	} else if num <= twoByteMax {
		return encodeTwoByte(uint16(num)), nil
	} else if num <= fourByteMax {
		return encodeFourByte(uint32(num)), nil
	} else if num <= eightByteMax {
		return encodeEightByte(uint64(num)), nil
	}
	return nil, unknownErr
}

func encodeOneByte(num uint8) []byte {
	return []byte{num}
}

func encodeTwoByte(num uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, num)
	b[0] = b[0] | 0x40
	return b
}

func encodeFourByte(num uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, num)
	b[0] = b[0] | 0x80
	return b
}

func encodeEightByte(num uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, num)
	b[0] = b[0] | 0xc0
	return b
}

func ToUint64(b []byte) (uint64, uint8) {
	length := 1 << (b[0] >> 6)
	b[0] = b[0] & 0x3f
	header := make([]byte, 8-length)
	return binary.BigEndian.Uint64(append(header, b[:length]...)), uint8(length)
}
