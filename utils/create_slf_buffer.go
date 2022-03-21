package utils

import "encoding/binary"

// NOTE: see SLF.md

// ensures the given buffer is of the given size,
// fills the remaining bytes with 0s if it doesn't
func EnsureBufferSize(buffer []byte, size int) []byte {
	var length int = len(buffer)

	var missingLength int = size - length

	for i := 0; i < missingLength; i++ {
		buffer = append(buffer, byte(0))
	}

	return buffer
}

// creates a basic slf file buffer
func CreateSlfBuffer(name string, path string, entries []entryInformation) []byte {
	var buffer []byte

	buffer = append(buffer, EnsureBufferSize([]byte(name), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, EnsureBufferSize([]byte(path), INT_BUFFER_STRING_LENGTH)...)

	var entriesCount int = len(entries)

	buffer = append(buffer, EnsureBufferSize([]byte{byte(entriesCount)}, INT_BUFFER_NUMBER_LENGTH)...)

	// ensure slf header size in order to write entries data in the right place
	buffer = EnsureBufferSize(buffer, INT_SLF_BUFFER_HEADER_LENGTH)

	var bufferEntriesInfos []byte

	var entryOffset = INT_SLF_BUFFER_HEADER_LENGTH

	for i := 0; i < entriesCount; i++ {
		var entry entryInformation = entries[i]

		buffer = append(buffer, entry.data...)

		var bufferEntryInfo []byte

		bufferEntryInfo = append(bufferEntryInfo, EnsureBufferSize([]byte(entry.name), INT_BUFFER_STRING_LENGTH)...)

		var entryDataSize int = len(entry.data)

		var entryOffsetBufferPart []byte = make([]byte, INT_BUFFER_NUMBER_LENGTH)
		binary.LittleEndian.PutUint32(entryOffsetBufferPart, uint32(entryOffset))

		var entryDataLengthBufferPart []byte = make([]byte, INT_BUFFER_NUMBER_LENGTH)
		binary.LittleEndian.PutUint32(entryDataLengthBufferPart, uint32(entryDataSize))

		bufferEntryInfo = append(bufferEntryInfo, entryOffsetBufferPart...)
		bufferEntryInfo = append(bufferEntryInfo, entryDataLengthBufferPart...)

		// ensure slf entry size in order to write the next entries data in the right place
		bufferEntryInfo = EnsureBufferSize(bufferEntryInfo, INT_SLF_BUFFER_ENTRY_LENGTH)

		bufferEntriesInfos = append(bufferEntriesInfos, bufferEntryInfo...)

		entryOffset += entryDataSize
	}

	buffer = append(buffer, bufferEntriesInfos...)

	return buffer
}
