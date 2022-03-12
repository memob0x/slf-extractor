import calcHrtimeSpan from './calc-hrtime-span';
import log from './log';
import readFileBuffer from './read-file-buffer';

const { hrtime } = process;

const readSlfBuffer = async (slfFilePath) => {
  if (slfFilePath) {
    log(`Process:\t Reading ${slfFilePath}\n`);
  }

  const readTimeStart = hrtime();

  let readResult = null;

  try {
    readResult = await readFileBuffer(
      slfFilePath,

      (size) => log(`File size:\t ${size}\n`),

      (p) => log(`Progress:\t ${p}%\r`),
    );
  } catch (e) {
    log(`${e.message}\n`);
  }

  log(`Report:\t File read in ${calcHrtimeSpan(readTimeStart)}\n`);

  return readResult;
};

export default readSlfBuffer;
