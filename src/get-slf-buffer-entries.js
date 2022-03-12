// thanks to https://github.com/ja2-stracciatella/ja2-stracciatella/blob/e17fbcb5268d6719d605f112424dc5f15d2e29a6/rust/stracciatella/src/file_formats/slf.rs

// Each entry represents a file.
// Numeric values are in little endian.
// Strings are '\0' terminated and have unused bytes zeroed.
// The paths are case-insensitive and use the '\\' character as a directory separator.
// Probably the special names for current directory "." and parent directory ".." are not supported.
// The header contains a library path, it is a path relative to the default directory (Data dir).
// Each entry contains a file path, it is a path relative to the library path.
// The encoding of the paths is unknown, but so far I've only seen ASCII.

// File Structure
// - header - 532 bytes, always at the start of the file
// - data - any size, contains the data of the entries
// - entries - 280 bytes per entry, always at the end of the file

// # Header Structure
// - 256 byte string with the library name
// - 256 byte string with the library path (empty or terminated by '\\', relative to Data dir)
// - 4 byte signed number with the total number of entries
// - 4 byte signed number with the total number of entries that have state FILE_OK 0x00
// - 2 byte unsigned number with name iSort (not used, only saw 0xFFFF, probably means it's sorted)
// - 2 byte unsigned number with name iVersion (not used, only saw 0x0200, probably means v2.0)
// - 1 byte unsigned number with name fContainsSubDirectories (not used, saw 0 and 1)
// - 3 byte padding (4 byte alignment)
// - 4 byte signed number with name iReserved (not used)

// # Entry Structure
// - 256 byte string with the file path (relative to the library path)
// - 4 byte unsigned number with the offset of the file data in the library file
// - 4 byte unsigned number with the length of the file data in the library file
// - 1 byte unsigned number with the state of the entry (saw FILE_OK 0x00 and FILE_OLD 0x01)
// - 1 byte unsigned number with name ubReserved (not used)
// - 2 byte padding (4 byte alignment)
// - 8 byte FILETIME (not used, from windows,
//   the number of 10^-7 seconds (100-nanosecond intervals) from 1 Jan 1601)
// - 2 byte unsigned number with name usReserved2 (not used)
// - 2 byte padding (4 byte alignment)

const INT_BUFFER_STRING_LENGTH = 256;
const INT_BUFFER_NUMBER_LENGTH = 4;

const INT_SLF_BUFFER_ENTRY_LENGTH = 280;
const INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT = INT_BUFFER_STRING_LENGTH
  // lib name + lib path
  * 2;

const getSlfBufferEntriesCount = (buffer) => buffer.readInt32LE(
  INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT,

  INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT + INT_BUFFER_NUMBER_LENGTH,
);

const getSlfBufferEntriesBuffer = (buffer, entriesCount) => {
  const { length } = buffer;

  return buffer.slice(
    length - INT_SLF_BUFFER_ENTRY_LENGTH * entriesCount,

    length,
  );
};

const getSlfBufferEntriesInfos = (buffer) => {
  const entriesCount = getSlfBufferEntriesCount(buffer);

  const bufferEntries = getSlfBufferEntriesBuffer(buffer, entriesCount);

  const infos = [];

  for (let entryIndex = 0; entryIndex < entriesCount; entryIndex += 1) {
    infos.push(
      [
        bufferEntries.toString(
          'utf-8',

          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex,

          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex + INT_BUFFER_STRING_LENGTH,
        ),

        bufferEntries.readUInt16LE(
          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex + INT_BUFFER_STRING_LENGTH,

          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex + INT_BUFFER_STRING_LENGTH + INT_BUFFER_NUMBER_LENGTH,
        ),

        bufferEntries.readUInt16LE(
          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex + INT_BUFFER_STRING_LENGTH + INT_BUFFER_NUMBER_LENGTH,

          INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex + INT_BUFFER_STRING_LENGTH + INT_BUFFER_NUMBER_LENGTH + INT_BUFFER_NUMBER_LENGTH,
        ),
      ],
    );
  }

  return infos;
};

const sanitizeName = (name) => name.replace(/\W+/g, '');

const sanitizeFilename = (name) => {
  const extension = name.split('.').slice(0).pop();

  return `${sanitizeName(name.replace(extension, ''))}.${sanitizeName(extension)}`;
};

const getSlfBufferEntries = (buffer) => {
  if (!Buffer.isBuffer(buffer)) {
    throw new Error('Input file is not valid buffer.');
  }

  // library original name buffer.toString('utf-8', 0, 256);
  // library original path buffer.toString('utf-8', 256, 256 + 256));

  const entriesInfos = getSlfBufferEntriesInfos(buffer);

  const entriesData = entriesInfos.map((x) => [
    sanitizeFilename(x[0]),

    buffer.slice(x[1], x[2]),
  ]);

  return entriesData;
};

export default getSlfBufferEntries;
