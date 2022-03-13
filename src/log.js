const { stdout } = process;

const { write } = stdout;

/**
 * @param {string} value - The string value to be logged to command line.
 * @param {string} end - The endline string.
 * @returns {boolean} Whatever is retourned by the "process.stdout.write" function.
 */
const log = (value, end = '\n') => write.call(stdout, `${value}${end}`);

export default log;
