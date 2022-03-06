import { writeFile } from 'fs/promises';

const writeSlfBuffer = async (readResult, destinationPath) => {
  // TODO: proper extraction
  await writeFile(`${destinationPath}/output.txt`, readResult);
};

export default writeSlfBuffer;
