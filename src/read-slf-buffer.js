import calcHrtimeSpan from './calc-hrtime-span';
import formatBytes from './format-bytes';
import formatMilliseconds from './format-milliseconds';
import getSlfHeader from './get-slf-header';
import log from './log';
import readFileBuffer from './read-file-buffer';

const { hrtime } = process;

const onStat = (size) => log(`File is ok (${formatBytes(size)})`);

const onProgress = (percentage) => log(`Progress: ${percentage}%`, '\r');

/**
 * @param {string} slfFilePath - The path to the slf file to be read.
 * @returns {Promise<Buffer>} The read slf file buffer.
 */
const readSlfBuffer = async (slfFilePath) => {
  if (slfFilePath) {
    log(`Reading ${slfFilePath}`);
  }

  const readTimeStart = hrtime();

  try {
    const buffer = await readFileBuffer(
      slfFilePath,

      {
        onStat,

        onProgress,
      },
    );

    const [
      originalSlfName,

      originalSlfPath,
    ] = getSlfHeader(buffer);

    log(`Original file name: ${originalSlfName}`);
    log(`Original file path: ${originalSlfPath} (relative to the installation path "Data" folder)`);

    log(`File read in ${formatMilliseconds(calcHrtimeSpan(readTimeStart))}`);

    return buffer;
  } catch (e) {
    log(`Error: ${e.message}`);

    return null;
  }
};

export default readSlfBuffer;
