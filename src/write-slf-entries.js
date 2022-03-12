import { mkdir, writeFile } from 'fs/promises';
import { hrtime } from 'process';

import calcHrtimeSpan from './calc-hrtime-span';
import isExistentFolder from './is-existent-folder';
import log from './log';

const writeSlfEntries = async (entries, destinationPath) => {
  const extractionTimeStart = hrtime();

  log(`Process:\t Extracting to ${destinationPath}\n`);

  if (!await isExistentFolder(destinationPath)) {
    await mkdir(destinationPath);

    log(`Notice:\t ${destinationPath} folder wasn't there, so it has been created\n`);
  }

  try {
    await Promise.all(entries.map(([name, data]) => {
      const path = `${destinationPath}/${name}`;

      log(`Process:\t ${path} file\n`);

      return writeFile(path, data);
    }));
  } catch (e) {
    log(`${e.message}\n`);

    return;
  }

  log(`Report:\t File extracted in ${calcHrtimeSpan(extractionTimeStart)}\n`);
};

export default writeSlfEntries;
