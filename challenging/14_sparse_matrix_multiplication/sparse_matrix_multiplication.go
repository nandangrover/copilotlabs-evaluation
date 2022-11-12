package main

import (
	"fmt"
	"math"
	"sort"
)

func sparseMatrixMultiplication(matrixA [][]int, matrixB [][]int) [][]int {
    if len(matrixA[0]) != len(matrixB) {
        return [][]int{}
    }

    sparseA := getDictOfNonzeroCells(matrixA)
    sparseB := getDictOfNonzeroCells(matrixB)

    matrixC := make([][]int, len(matrixA))
    for i := range matrixC {
        matrixC[i] = make([]int, len(matrixB[0]))
    }

    for i, k := range sparseA.keys() {
        for j := range matrixB[0] {
            if (k, j) in sparseB.keys() {
                matrixC[i][j] += sparseA[(i, k)] * sparseB[(k, j)]
            }
        }
    }
    return matrixC
}


func getDictOfNonzeroCells(matrix [][]int) map[int]int {
    dictOfNonzeroCells := make(map[int]int)
    for i := range matrix {
        for j := range matrix[0] {
            if matrix[i][j] != 0 {
                dictOfNonzeroCells[(i, j)] = matrix[i][j]
            }
        }
    }
    return dictOfNonzeroCells
}

func main() {


matrixA := [][]int{
    {0, 2, 0},
    {0, -3, 5},
}
matrixB := [][]int{
    {0, 10, 0},
    {0, 0, 0},
    {0, 0, 4},
}

fmt.Println(sparseMatrixMultiplication(matrixA, matrixB))
}