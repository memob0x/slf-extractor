const { hrtime } = process;

/**
 * @param {number[]} start - The time span start "process.hrtime" function result.
 * @returns {number} The resulting time span in milliseconds.
 */
const calcHrtimeSpan = (start) => {
  const [s, ms] = hrtime(start);

  return s * 1000 + ms / 1000000;
};

export default calcHrtimeSpan;
