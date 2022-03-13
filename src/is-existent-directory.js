import { stat } from 'fs/promises';

/**
 * @param {string} path - The given path to be checked.
 * @returns {Promise<boolean>} Whether the given path leads to an existent folder or not,
 * false if not.
 */
const isExistentDirectory = async (path) => {
  if (!path) {
    return false;
  }

  try {
    const folderStats = await stat(path);

    return folderStats.isDirectory();
  } catch (_e) {
    return false;
  }
};

export default isExistentDirectory;
