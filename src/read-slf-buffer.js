import calcHrtimeSpan from './calc-hrtime-span';
import log from './log';
import readFileBuffer from './read-file-buffer';

const { hrtime } = process;

const onStat = (size) => log(`File size:\t ${size}\n`);

const onProgress = (percentage) => log(`Progress:\t ${percentage}%\r`);

const readSlfBuffer = async (slfFilePath) => {
  if (slfFilePath) {
    log(`Process:\t Reading ${slfFilePath}\n`);
  }

  const readTimeStart = hrtime();

  let readResult = null;

  try {
    readResult = await readFileBuffer(
      slfFilePath,

      {
        onStat,

        onProgress,
      },
    );
  } catch (e) {
    log(`${e.message}\n`);
  }

  log(`Report:\t File read in ${calcHrtimeSpan(readTimeStart)}\n`);

  return readResult;
};

export default readSlfBuffer;
