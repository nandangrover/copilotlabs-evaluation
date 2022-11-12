const  numberOfWaysToTraverseGraph = require('./traverse_graph');

test("Test Case #1", () => {
  const width = 4;
  const height = 3;
  const expected = 10;
  const actual = numberOfWaysToTraverseGraph(width, height);
  expect(actual).toEqual(expected);
});
