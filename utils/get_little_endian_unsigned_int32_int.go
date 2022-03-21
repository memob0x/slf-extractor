package utils

import "encoding/binary"

func GetLittleEndianUnsignedInt32Int(buffer []byte, sliceStart int, sliceEnd int) int {
	return int(binary.LittleEndian.Uint32(buffer[sliceStart:sliceEnd]))
}
