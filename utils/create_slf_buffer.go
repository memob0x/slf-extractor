package utils

// NOTE: see SLF.md

// TODO: check, maybe make() with the right arguments could have been enough...
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

func CreateSlfBufferHeader(name string, path string, entriesCount int) []byte {
	var buffer []byte

	buffer = append(buffer, EnsureBufferSize([]byte(name), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, EnsureBufferSize([]byte(path), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, EnsureBufferSize([]byte{byte(entriesCount)}, INT_BUFFER_NUMBER_LENGTH)...)

	// ensure slf header size in order to write entries data in the right place
	buffer = EnsureBufferSize(buffer, INT_SLF_BUFFER_HEADER_LENGTH)

	return buffer
}

func CreateSlfBufferEntry(name string, offset int, length int) []byte {
	var buffer []byte

	buffer = append(buffer, EnsureBufferSize([]byte(name), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, GetLittleEndianUnsignedInt32Buffer(offset, INT_BUFFER_NUMBER_LENGTH)...)

	buffer = append(buffer, GetLittleEndianUnsignedInt32Buffer(length, INT_BUFFER_NUMBER_LENGTH)...)

	// ensure slf entry size in order to write the next entries data in the right place
	buffer = EnsureBufferSize(buffer, INT_SLF_BUFFER_ENTRY_LENGTH)

	return buffer
}

// creates a basic slf file buffer
func CreateSlfBuffer(name string, path string, entries []entryInformation) []byte {
	var entriesCount int = len(entries)

	var buffer []byte = CreateSlfBufferHeader(name, path, entriesCount)

	var bufferEntriesInfos []byte

	var entryOffset = INT_SLF_BUFFER_HEADER_LENGTH

	for i := 0; i < entriesCount; i++ {
		var entry entryInformation = entries[i]

		buffer = append(buffer, entry.data...)

		var entryDataSize int = len(entry.data)

		bufferEntriesInfos = append(bufferEntriesInfos, CreateSlfBufferEntry(entry.name, entryOffset, entryDataSize)...)

		entryOffset += entryDataSize
	}

	buffer = append(buffer, bufferEntriesInfos...)

	return buffer
}
