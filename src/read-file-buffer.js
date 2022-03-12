import { stat } from 'fs/promises';
import { createReadStream } from 'fs';

const readStreamOptions = {
  // TODO: check what is this all about...

  // flag: 'a+',
  // encoding: 'ASCII',
  // start: 5,
  // end: 64,

  // highWaterMark: 16,
};

const readFileBufferWithStreams = (filePath, fileSize, onProgress) => (
  new Promise((resolve, reject) => {
    const readStream = createReadStream(
      filePath,

      readStreamOptions,
    );

    const readChunks = [];

    let readLength = 0;

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
const readFileBuffer = async (filePath, onStat, onProgress) => {
  if (!filePath) {
    throw new Error('Missing input file.');
  }

  const stats = await stat(filePath);

  if (!stats) {
    throw new Error(`${filePath} is either invalid or unexistent.`);
  }

  const fileSize = stats.size;

  if (onStat) {
    onStat(fileSize);
  }

  return readFileBufferWithStreams(filePath, fileSize, onProgress);
};

export default readFileBuffer;
