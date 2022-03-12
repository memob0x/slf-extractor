import { stat } from 'fs/promises';

const isExistentFolder = async (folderPath) => {
  if (!folderPath) {
    return false;
  }

  try {
    const folderStats = await stat(folderPath);

    return folderStats.isDirectory();
  } catch (_e) {
    return false;
  }
};

export default isExistentFolder;
