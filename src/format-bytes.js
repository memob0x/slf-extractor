const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];

/**
 * @param {number} bytes - The bytes value to be formatted.
 * @returns {string} The given bytes formatted with unit values.
 */
const formatBytes = (bytes) => {
  if (bytes === 0) {
    return 'n/a';
  }

  const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10);

  if (i === 0) {
    return `${bytes} ${sizes[i]}`;
  }

  return `${(bytes / (1024 ** i)).toFixed(1)} ${sizes[i]}`;
};

export default formatBytes;
