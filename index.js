import path from 'path';
import { argv } from 'process';

import extractSlf from './extract-slf';

const [
  slfFilePath,

  extractionDestinationPath = './output',
] = argv.slice(2);

extractSlf(
  path.resolve(slfFilePath),

  path.resolve(extractionDestinationPath),
);
