const { stdout } = process;

const log = (...args) => stdout.write(...args);

export default log;
