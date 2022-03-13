import { mkdir, writeFile } from 'fs/promises';
import { hrtime } from 'process';

import calcHrtimeSpan from './calc-hrtime-span';
import formatBytes from './format-bytes';
import isExistentDirectory from './is-existent-directory';
import log from './log';

import {
  STRING_ENCODING,
} from './constants';
import formatMilliseconds from './format-milliseconds';

const getWriteSlfEntriesReducer = (destinationPath) => (
  function writeSlfEntriesReducer([totalLength, writings], [name, data]) {
    const path = `${destinationPath}/${name}`;

    const { length } = data;

    log(`Extracting ${path} file (${formatBytes(length)})`);

    writings.push(writeFile(path, data, STRING_ENCODING));

    return [
      totalLength + length,

      writings,
    ];
  }
);

/**
 * @param {any[]} entries - The slf file entries informations,
 * every array item is composed by a string with the name of the file
 * and two integeres with the slice arguments.
 * @param {string} destinationPath - The extraction output destination path,
 * whete the extracted files needs to be saved.
 * @returns {Promise<boolean>} Whether the writing process has been successful or not.
 */
const writeSlfEntries = async (entries, destinationPath) => {
  if (!Array.isArray(entries)) {
    log('Error: Invalid entries provided to writing function');

    return false;
  }

  const extractionTimeStart = hrtime();

  log(`Extracting to ${destinationPath}`);

  if (!await isExistentDirectory(destinationPath)) {
    await mkdir(destinationPath);

    log(`The directory ${destinationPath} wasn't there, so it has been created`);
  }

  try {
    const [
      totalLength,

      writings,
    ] = entries.reduce(
      getWriteSlfEntriesReducer(destinationPath),

      [
        0,

        [],
      ],
    );

    await Promise.all(writings);

    log(`All file contents have been extracted (${formatBytes(totalLength)}) in ${formatMilliseconds(calcHrtimeSpan(extractionTimeStart))}`);

    return true;
  } catch (e) {
    log(`Error: ${e.message}`);

    return false;
  }
};

export default writeSlfEntries;
