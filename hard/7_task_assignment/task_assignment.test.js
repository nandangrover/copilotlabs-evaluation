const  taskAssignment = require('./task_assignment');
// const assert = require('assert')

test("Test Case #1", () => {
  const k = 3;
  const tasks = [1, 3, 5, 3, 1, 4];
  const expected = [
    [4, 2],
    [0, 5],
    [3, 1],
  ];
  const actual = taskAssignment(k, tasks);
  expect(actual).toEqual(expected);
});
