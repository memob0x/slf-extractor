import runtimeReadSlfBuffer from './runtime-read-slf-buffer';
import runtimeExtractSlfBuffer from './runtime-write-slf-buffer';

const extractSlf = async (slfFilePath, extractionDestinationPath) => runtimeExtractSlfBuffer(
  await runtimeReadSlfBuffer(slfFilePath),

  extractionDestinationPath,
);

export default extractSlf;
