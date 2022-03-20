package main

import "encoding/binary"

func GetLittleEndianUnsignedInt32(buffer []byte, sliceStart int, sliceEnd int) int {
	return int(binary.LittleEndian.Uint32(buffer[sliceStart:sliceEnd]))
}
