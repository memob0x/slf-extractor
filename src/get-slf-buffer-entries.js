// NOTE: see SLF.md

import isBuffer from './is-buffer';
import sanitizeStringFilename from './sanitize-string-filename';

import {
  INT_BUFFER_NUMBER_LENGTH,
  INT_BUFFER_STRING_LENGTH,
  INT_SLF_BUFFER_ENTRY_LENGTH,
  INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT,
  STRING_ENCODING,
} from './constants';

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

const getSlfBufferEntryInfos = (bufferEntries, entryIndex) => {
  const info = [];

  let pointer0 = INT_SLF_BUFFER_ENTRY_LENGTH * entryIndex;

  let pointer1 = pointer0 + INT_BUFFER_STRING_LENGTH;

  const name = bufferEntries.toString(
    STRING_ENCODING,

    pointer0,

    pointer1,
  );

  info.push(sanitizeStringFilename(name));

  pointer0 = pointer1;

  pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH;

  const dataSliceStart = bufferEntries.readUInt32LE(
    pointer0,

    pointer1,
  );

  info.push(dataSliceStart);

  pointer0 = pointer1;

  pointer1 = pointer0 + INT_BUFFER_NUMBER_LENGTH;

  const dataSliceLength = bufferEntries.readUInt32LE(
    pointer0,

    pointer1,
  );

  info.push(dataSliceLength);

  return info;
};

const getSlfBufferEntriesInfos = (buffer) => {
  const entriesCount = getSlfBufferEntriesCount(buffer);

  const bufferEntries = getSlfBufferEntriesBuffer(buffer, entriesCount);

  const infos = [];

  for (let entryIndex = 0; entryIndex < entriesCount; entryIndex += 1) {
    infos.push(getSlfBufferEntryInfos(bufferEntries, entryIndex));
  }

  return infos;
};

/**
 * @throws Will throw an error if the argument is not a valid buffer object.
 * @param {Buffer} buffer - The read slf file buffer.
 * @returns {any[]} The slf file entries informations,
 * every array item is composed by a string with the name of the file
 * and two integeres with the slice arguments.
 */
const getSlfBufferEntries = (buffer) => {
  if (!isBuffer(buffer)) {
    throw new Error('Input file is not valid buffer.');
  }

  return getSlfBufferEntriesInfos(buffer)
    .map(([name, offset, length]) => [
      name,

      buffer.slice(
        offset,

        offset + length,
      ),
    ]);
};

export default getSlfBufferEntries;
