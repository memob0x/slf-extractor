import getSlfBufferEntries from './get-slf-buffer-entries';
import isBuffer from './is-buffer';
import readSlfBuffer from './read-slf-buffer';
import writeSlfEntries from './write-slf-entries';

/**
 * @param {string} slfFilePath - The path to the slf file to be read.
 * @param {string} destinationPath - The extraction output destination path,
 * whete the extracted files needs to be saved.
 * @returns {Promise<boolean>} Whether the writing process has been successful or not.
 */
const extractSlfContents = async (slfFilePath, destinationPath) => {
  const buffer = await readSlfBuffer(slfFilePath);

  if (!isBuffer(buffer)) {
    return null;
  }

  return writeSlfEntries(
    getSlfBufferEntries(buffer),

    destinationPath,
  );
};

export default extractSlfContents;
