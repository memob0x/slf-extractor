import { stat } from 'fs/promises';
import { createReadStream } from 'fs';

const readFileBufferWithStreams = (filePath, fileSize, options) => (
  new Promise((resolve, reject) => {
    const readStream = createReadStream(
      filePath,

      options,
    );

    const readChunks = [];

    let readLength = 0;

    const { onProgress } = options || {};

    readStream.on('data', (chunk) => {
      readChunks.push(chunk);

      readLength += chunk.length;

      if (onProgress) {
      // eslint-disable-next-line no-mixed-operators
        const percentage = (readLength / fileSize * 100).toFixed(2);

        onProgress(percentage);
      }
    });

    readStream.on('end', () => resolve(Buffer.concat(readChunks)));

    readStream.on('error', reject);
  })
);

/**
 * @param {string} filePath - The path to the file to be read.
 * @param {object} options - The options object to be passed to "fs.createReadStream" function
 * along with "onStat" and "onProgress" callbacks.
 * @returns {Promise<Buffer>} The read file buffer.
 */
const readFileBuffer = async (filePath, options) => {
  if (!filePath) {
    throw new Error('Missing input file.');
  }

  const { size } = await stat(filePath) || {};

  if (!size) {
    throw new Error(`${filePath} is either invalid or unexistent.`);
  }

  const { onStat } = options || {};

  if (onStat) {
    onStat(size);
  }

  return readFileBufferWithStreams(filePath, size, options);
};

export default readFileBuffer;
