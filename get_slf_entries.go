package main

// NOTE: see SLF.md

type entryInformation struct {
	name string

	data []byte
}

func GetSlfBufferEntryInfos(bufferAreaEntriesInfos []byte, entryIndex int, buffer []byte) entryInformation {
	var info entryInformation = entryInformation{}

	var pointer0 int = INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex

	var pointer1 int = pointer0 + INT_BUFFER_STRING_LENGTH

	info.name = SanitizeStringFilename(string(bufferAreaEntriesInfos[pointer0:pointer1]))

	pointer0 = pointer1

	pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH

	var dataSliceStart int = GetLittleEndianUnsignedInt32(bufferAreaEntriesInfos, pointer0, pointer1)

	pointer0 = pointer1

	pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH

	var dataSliceLength int = GetLittleEndianUnsignedInt32(bufferAreaEntriesInfos, pointer0, pointer1)

	info.data = buffer[dataSliceStart : dataSliceStart+dataSliceLength]

	return info
}

// Gets the slf file entries informations from a given slf file buffer
func GetSlfBufferEntries(buffer []byte) []entryInformation {
	var infos []entryInformation = []entryInformation{}

	var entriesCount int = GetLittleEndianUnsignedInt32(
		buffer,

		INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT,

		INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT+INT_BUFFER_NUMBER_LENGTH,
	)

	var bufferLength int = len(buffer)

	var bufferAreaEntriesInfos []byte = buffer[bufferLength-INT_SLF_BUFFER_ENTRY_LENGTH*entriesCount : bufferLength]

	for entryIndex := 0; entryIndex < entriesCount; entryIndex++ {
		infos = append(infos, GetSlfBufferEntryInfos(bufferAreaEntriesInfos, entryIndex, buffer))
	}

	return infos
}
