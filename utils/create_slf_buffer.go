package utils

// NOTE: see SLF.md

// TODO: check, maybe make() with the right arguments could have been enough...
// ensures the given buffer is of the given size,
// fills the remaining bytes with 0s if it doesn't
func ensureBufferSize(buffer []byte, size int) []byte {
	var length int = len(buffer)

	var missingLength int = size - length

	for i := 0; i < missingLength; i++ {
		buffer = append(buffer, byte(0))
	}

	return buffer
}

func createSlfBufferHeader(name string, path string, entriesCount int) []byte {
	var buffer []byte

	buffer = append(buffer, ensureBufferSize([]byte(name), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, ensureBufferSize([]byte(path), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, ensureBufferSize([]byte{byte(entriesCount)}, INT_BUFFER_NUMBER_LENGTH)...)

	// ensure slf header size in order to write entries data in the right place
	buffer = ensureBufferSize(buffer, INT_SLF_BUFFER_HEADER_LENGTH)

	return buffer
}

func createSlfBufferEntry(name string, offset int, length int) []byte {
	var buffer []byte

	buffer = append(buffer, ensureBufferSize([]byte(name), INT_BUFFER_STRING_LENGTH)...)

	buffer = append(buffer, GetLittleEndianUnsignedInt32Buffer(offset, INT_BUFFER_NUMBER_LENGTH)...)

	buffer = append(buffer, GetLittleEndianUnsignedInt32Buffer(length, INT_BUFFER_NUMBER_LENGTH)...)

	// ensure slf entry size in order to write the next entries data in the right place
	buffer = ensureBufferSize(buffer, INT_SLF_BUFFER_ENTRY_LENGTH)

	return buffer
}

// creates a basic slf file buffer
func CreateSlfBuffer(name string, path string, entries []SlfEntry) []byte {
	var entriesCount int = len(entries)

	var buffer []byte = createSlfBufferHeader(name, path, entriesCount)

	var bufferEntriesInfos []byte

	var entryOffset = INT_SLF_BUFFER_HEADER_LENGTH

	for i := 0; i < entriesCount; i++ {
		var entry SlfEntry = entries[i]

		buffer = append(buffer, entry.Data...)

		var entryDataSize int = len(entry.Data)

		bufferEntriesInfos = append(bufferEntriesInfos, createSlfBufferEntry(entry.Name, entryOffset, entryDataSize)...)

		entryOffset += entryDataSize
	}

	buffer = append(buffer, bufferEntriesInfos...)

	return buffer
}
