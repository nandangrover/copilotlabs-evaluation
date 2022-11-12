const sparse_matrix_multiplication = require('./sparse_matrix_multiplication');

describe("sparse matrix mult", () => {
  

  test('Test case 1', () => {
      const matrix_a = [
          [0, 2, 0],
          [0, -3, 5],
      ];
      const matrix_b = [
          [0, 10, 0],
          [0, 0, 0],
          [0, 0, 4],
      ];
      const expected = [
          [0, 0, 0],
          [0, 0, 20],
      ];
      const actual = sparse_matrix_multiplication(matrix_a, matrix_b);
      expect(actual).toEqual(expected);
  });
});
