// NOTE: see SLF.md

import isBuffer from './is-buffer';

import sanitizeStringFilename from './sanitize-string-filename';

import {
  INT_BUFFER_STRING_LENGTH,
  STRING_ENCODING,
} from './constants';

/**
 * @throws Will throw when the given argument is either not a valid buffer
 * or a buffer of an invalid slf file.
 * @param {Buffer} buffer - The slf file buffer.
 * @returns {string[]} The main header informations: the original slf file name and path.
 */
const getSlfHeader = (buffer) => {
  if (!isBuffer(buffer)) {
    throw new Error('Input file is not valid buffer.');
  }

  const header = [];

  let pointer0 = 0;
  let pointer1 = INT_BUFFER_STRING_LENGTH;

  const name = sanitizeStringFilename(buffer.toString(STRING_ENCODING, pointer0, pointer1));

  if (!name.toLowerCase().endsWith('slf')) {
    throw new Error('Not a valid slf file');
  }

  header.push(name);

  pointer0 = pointer1;
  pointer1 = pointer0 + INT_BUFFER_STRING_LENGTH;

  const path = buffer.toString(STRING_ENCODING, pointer0, pointer1).replaceAll('\\', '/');

  header.push(path);

  return header;
};

export default getSlfHeader;
