import sanitizeStringWord from './sanitize-string-word';

/**
 * @param {string} fileName - The given file name string to be sanitized.
 * @returns {string} The given file name with all invalid characters removed.
 */
const sanitizeStringFilename = (fileName) => {
  const extension = fileName.split('.').pop();

  return `${sanitizeStringWord(fileName.replace(extension, ''))}.${sanitizeStringWord(extension)}`;
};

export default sanitizeStringFilename;
