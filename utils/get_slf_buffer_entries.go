package utils

// NOTE: see SLF.md

type SlfEntry struct {
	Name string

	Data []byte
}

func getSlfBufferEntryInfos(bufferAreaEntriesInfos []byte, entryIndex int, buffer []byte) SlfEntry {
	var info SlfEntry = SlfEntry{}

	var pointer0 int = INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex

	var pointer1 int = pointer0 + INT_BUFFER_STRING_LENGTH

	info.Name = SanitizeStringFilename(string(bufferAreaEntriesInfos[pointer0:pointer1]))

	pointer0 = pointer1

	pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH

	var dataSliceStart int = GetLittleEndianUnsignedInt32Int(bufferAreaEntriesInfos, pointer0, pointer1)

	pointer0 = pointer1

	pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH

	var dataSliceLength int = GetLittleEndianUnsignedInt32Int(bufferAreaEntriesInfos, pointer0, pointer1)

	info.Data = buffer[dataSliceStart : dataSliceStart+dataSliceLength]

	return info
}

// Gets the slf file entries informations from a given slf file buffer
func GetSlfBufferEntries(buffer []byte) []SlfEntry {
	var infos []SlfEntry = []SlfEntry{}

	var entriesCount int = GetLittleEndianUnsignedInt32Int(
		buffer,

		INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT,

		INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT+INT_BUFFER_NUMBER_LENGTH,
	)

	var bufferLength int = len(buffer)

	var bufferAreaEntriesInfos []byte = buffer[bufferLength-INT_SLF_BUFFER_ENTRY_LENGTH*entriesCount : bufferLength]

	for entryIndex := 0; entryIndex < entriesCount; entryIndex++ {
		infos = append(infos, getSlfBufferEntryInfos(bufferAreaEntriesInfos, entryIndex, buffer))
	}

	return infos
}
