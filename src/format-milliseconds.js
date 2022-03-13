/**
 * @param {number} milliseconds - The given milliseconds to be formatted.
 * @returns {string} The given milliseconds formatted in seconds, minutes etc...
 */
const formatMilliseconds = (milliseconds) => {
  if (milliseconds < 1000) {
    return `${milliseconds.toFixed(1)} milliseconds`;
  }

  const seconds = (milliseconds / 1000).toFixed(1);

  if (seconds < 60) {
    return `${seconds} seconds`;
  }

  const minutes = (milliseconds / (1000 * 60)).toFixed(1);

  if (minutes < 60) {
    return `${minutes} ninutes`;
  }

  const hours = (milliseconds / (1000 * 60 * 60)).toFixed(1);

  if (hours < 24) {
    return `${hours} hours`;
  }

  const days = (milliseconds / (1000 * 60 * 60 * 24)).toFixed(1);

  return `${days} days`;
};

export default formatMilliseconds;
