import { mkdir } from 'fs/promises';
import { hrtime } from 'process';

import isExistentFolder from '../is-existent-folder';
import writeSlfBuffer from '../write-slf-buffer';
import log from '../log';
import calcHrtimeSpan from '../calc-hrtime-span';

const runtimeReadSlfBuffer = async (readResult, extractionDestinationPath) => {
  const extractionTimeStart = hrtime();

  log(`Process:\t Extracting to ${extractionDestinationPath}\n`);

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

  log(`Report:\t File extracted in ${calcHrtimeSpan(extractionTimeStart)}\n`);
};

export default runtimeReadSlfBuffer;
