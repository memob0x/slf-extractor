const { hrtime } = process;

const calcHrtimeSpan = (start) => hrtime(start)[1] / 1000000;

export default calcHrtimeSpan;
