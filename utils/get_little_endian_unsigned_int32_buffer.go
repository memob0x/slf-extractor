package utils

import "encoding/binary"

func GetLittleEndianUnsignedInt32Buffer(value int, size int) []byte {
	var entryDataLengthBufferPart []byte = make([]byte, size)

	binary.LittleEndian.PutUint32(entryDataLengthBufferPart, uint32(value))

	return entryDataLengthBufferPart
}
