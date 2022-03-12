const { stdout } = process;

const { write: log } = stdout;

export default log.bind(stdout);
