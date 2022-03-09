import path from 'path';
import { argv } from 'process';

import runtimeExtractSlf from './src/utils/runtime/runtime-extract-slf';

const [
  slfFilePath,

  extractionDestinationPath = './output',
] = argv.slice(2);

runtimeExtractSlf(
  path.resolve(slfFilePath),

  path.resolve(extractionDestinationPath),
);
