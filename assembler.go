package sacio

import (
	"encoding/binary"
	"math"
)

// assembleString assembles a string from 4 bytes
func assembleString(data []byte) K {
	return K(data)
}

// assembleBool assembles a bool from 4 bytes
func assembleBool(data []byte, bitOrder int) L {
	if bitOrder == LSBFIRST {
		return data[0] != 0x00
	}

	return data[3] != 0x00
}

// assembleInt32 assembles a int32 from 4 bytes
func assembleInt32(data []byte, bitOrder int) N {
	var bits uint32
	if bitOrder == LSBFIRST {
		bits = uint32(data[3])<<24 | uint32(data[2])<<16 | uint32(data[1])<<8 | uint32(data[0])
		return N(bits)
	}

	bits = uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	return N(bits)
}

// assembleFloat32 assembles a float32 from 4 bytes
func assembleFloat32(data []byte, bitOrder int) F {
	var bits uint32
	if bitOrder == LSBFIRST {
		bits = uint32(data[3])<<24 | uint32(data[2])<<16 | uint32(data[1])<<8 | uint32(data[0])
		return F(math.Float32frombits(bits))
	}

	bits = uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	return F(math.Float32frombits(bits))
}

// assembleEnum assembles enums to string from 4 bytes
func assembleEnum(data []byte, bitOrder int) I {
	bits := assembleInt32(data, bitOrder)
	for _, v := range ENUMS {
		if v.Index == int(bits) {
			return I(v.Value)
		}
	}

	return ""
}

// disassembleFloat32 disassembles a float32 to 4 bytes
func disassembleFloat32(data F, bitOrder int) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, math.Float32bits(float32(data)))

	if bitOrder == LSBFIRST {
		for i := 0; i < 2; i++ {
			bytes[i], bytes[3-i] = bytes[3-i], bytes[i]
		}
	}

	return bytes
}

// disassembleInt32 disassembles a int32 to 4 bytes
func disassembleInt32(data N, bitOrder int) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(data))

	if bitOrder == LSBFIRST {
		for i := 0; i < 2; i++ {
			bytes[i], bytes[3-i] = bytes[3-i], bytes[i]
		}
	}

	return bytes
}

// disassembleBool disassembles a bool to 4 bytes
func disassembleBool(data L, bitOrder int) []byte {
	if data && bitOrder == LSBFIRST {
		return []byte{0, 0, 0, 1}
	} else if data && bitOrder == MSBFIRST {
		return []byte{1, 0, 0, 0}
	}

	return []byte{0, 0, 0, 0}
}

// disassembleString disassembles a string to 4 bytes
func disassembleString(data K) []byte {
	bytes := []byte(data)
	if len(bytes) > 8 {
		bytes = bytes[:8]
	}

	return bytes
}

// disassembleEnum disassembles a enum to 4 bytes
func disassembleEnum(data I, bitOrder int) []byte {
	for _, v := range ENUMS {
		if v.Value == string(data) {
			return disassembleInt32(N(v.Index), bitOrder)
		}
	}

	return disassembleInt32(0, bitOrder)
}
