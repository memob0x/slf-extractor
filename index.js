import path from 'path';
import { argv } from 'process';

import extractSlfEntries from './src/extract-slf-entries';

const [
  slfFilePath,

  extractionDestinationPath = './output',
] = argv.slice(2);

extractSlfEntries(
  path.resolve(slfFilePath),

  path.resolve(extractionDestinationPath),
);
