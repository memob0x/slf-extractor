import { mkdir } from 'fs/promises';

import readSlfBuffer from './read-slf-buffer';
import isExistentFolder from './is-existent-folder';
import writeSlfBuffer from './write-slf-buffer';

const { hrtime, stdout } = process;

const log = (...args) => stdout.write(...args);

const calcHrtimeSpan = (start) => hrtime(start)[1] / 1000000;

const extractSlf = async (slfFilePath, extractionDestinationPath) => {
  // READ
  // ------------------------------------------------------------

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

    return;
  }

  log(`Report:\t File read in ${calcHrtimeSpan(readTimeStart)}\n`);

  // EXTRACTION
  // ------------------------------------------------------------

  const extractionTimeStart = hrtime();

  log(`Process:\t Extracting ${slfFilePath} to ${extractionDestinationPath}\n`);

  if (!await isExistentFolder(extractionDestinationPath)) {
    await mkdir(extractionDestinationPath);

    log(`Notice:\t ${extractionDestinationPath} folder wan't there, so it has been created\n`);
  }

  try {
    await writeSlfBuffer(readResult, extractionDestinationPath);
  } catch (e) {
    log(`${e.message}\n`);

    return;
  }

  log(`Report:\t File extraction in ${calcHrtimeSpan(extractionTimeStart)}\n`);
};

export default extractSlf;
