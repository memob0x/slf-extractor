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
 * @param {string} filePath The path to the file to be read
 * @returns {Promise<any>}
 */
const readFileBuffer = async (filePath, options) => {
  if (!filePath) {
    throw new Error('Missing input file.');
  }

  const stats = await stat(filePath);

  if (!stats) {
    throw new Error(`${filePath} is either invalid or unexistent.`);
  }

  const fileSize = stats.size;

  const { onStat } = options || {};

  if (onStat) {
    onStat(fileSize);
  }

  return readFileBufferWithStreams(filePath, fileSize, options);
};

export default readFileBuffer;
