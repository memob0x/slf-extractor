import getSlfBufferEntries from './get-slf-buffer-entries';
import readSlfBuffer from './read-slf-buffer';
import writeSlfEntries from './write-slf-entries';

const extractSlfEntries = async (slfFilePath, destinationPath) => writeSlfEntries(
  getSlfBufferEntries(await readSlfBuffer(slfFilePath)),

  destinationPath,
);

export default extractSlfEntries;
