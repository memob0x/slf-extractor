import calcHrtimeSpan from '../calc-hrtime-span';
import log from '../log';
import readSlfBuffer from '../read-slf-buffer';

const { hrtime } = process;

const runtimeReadSlf = async (slfFilePath) => {
  if (slfFilePath) {
    log(`Process:\t Reading ${slfFilePath}\n`);
  }

  const readTimeStart = hrtime();

  let readResult = null;

  try {
    readResult = await readSlfBuffer(
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

export default runtimeReadSlf;
