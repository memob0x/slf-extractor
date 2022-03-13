/**
 * @param {string} string - Performs a global search for word characters on the given argument,
 * removing invalid ones.
 * @returns {string} The given argument without invalid word characters.
 */
const sanitizeStringWord = (string) => string.replace(/\W+/g, '');

export default sanitizeStringWord;
