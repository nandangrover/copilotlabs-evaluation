const twoEdgeConnectedGraph = require('./two_edge_connected_graph');

it('Test Case #1', function () {
  const input = [
    [1, 2, 5],
    [0, 2],
    [0, 1, 3],
    [2, 4, 5],
    [3, 5],
    [0, 3, 4],
  ];
  const expected = true;
  const actual = twoEdgeConnectedGraph(input);
  expect(actual).toEqual(expected);
});
