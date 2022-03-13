import path from 'path';
import { argv } from 'process';

import extractSlfContents from './src/extract-slf-contents';

const [
  slfFilePath,

  extractionDestinationPath = './output',
] = argv.slice(2);

extractSlfContents(
  path.resolve(slfFilePath),

  path.resolve(extractionDestinationPath),
);
