function sparse_matrix_multiplication(matrix_a, matrix_b) {
    if (matrix_a[0].length != matrix_b.length) {
        return [[]];
    }

    var sparse_a = get_dict_of_nonzero_cells(matrix_a);
    var sparse_b = get_dict_of_nonzero_cells(matrix_b);

    var matrix_c = [...Array(matrix_a.length)].map(e => Array(matrix_b[0].length).fill(0));

    for (var [i, k] of Object.keys(sparse_a)) {
        for (var j = 0; j < matrix_b[0].length; j++) {
            if ([k, j] in Object.keys(sparse_b)) {
                matrix_c[i][j] += sparse_a[[i, k]] * sparse_b[[k, j]];
            }
        }
    }
    return matrix_c;
}


function get_dict_of_nonzero_cells(matrix) {
    var dict_of_nonzero_cells = {};
    for (var i = 0; i < matrix.length; i++) {
        for (var j = 0; j < matrix[0].length; j++) {
            if (matrix[i][j] != 0) {
                dict_of_nonzero_cells[[i, j]] = matrix[i][j];
            }
        }
    }
    return dict_of_nonzero_cells;
}


var matrix_a = [
    [0, 2, 0],
    [0, -3, 5],
];
var matrix_b = [
    [0, 10, 0],
    [0, 0, 0],
    [0, 0, 4],
];

console.log(sparse_matrix_multiplication(matrix_a, matrix_b));

module.exports = sparse_matrix_multiplication;