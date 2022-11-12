package main

import "fmt"

func main() {
	fmt.Println(nonAttackingQueens(4)) // 2
	fmt.Println(nonAttackingQueens(5)) // 10
	fmt.Println(nonAttackingQueens(6)) // 4
}

func nonAttackingQueens(n int) int {
	blockedColumns := make(map[int]bool)
	blockedUpDiagonals := make(map[int]bool)
	blockedDownDiagonals := make(map[int]bool)
	return getNumberOfNonAttackingQueenPlacements(0, blockedColumns, blockedUpDiagonals, blockedDownDiagonals, n)
}

func getNumberOfNonAttackingQueenPlacements(row int, blockedColumns map[int]bool, blockedUpDiagonals map[int]bool, blockedDownDiagonals map[int]bool, boardSize int) int {
	if row == boardSize {
		return 1
	}

	validPlacements := 0
	for col := 0; col < boardSize; col++ {
		if isNonAttackingPlacement(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals) {
			placeQueen(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals)
			validPlacements += getNumberOfNonAttackingQueenPlacements(
				row+1, blockedColumns, blockedUpDiagonals, blockedDownDiagonals, boardSize,
			)
			removeQueen(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals)
		}
	}

	return validPlacements
}

// This is always an O(1)-time operation.
func isNonAttackingPlacement(row int, col int, blockedColumns map[int]bool, blockedUpDiagonals map[int]bool, blockedDownDiagonals map[int]bool) bool {
	if blockedColumns[col] {
		return false
	}
	if blockedUpDiagonals[row+col] {
		return false
	}
	if blockedDownDiagonals[row-col] {
		return false
	}

	return true
}

func placeQueen(row int, col int, blockedColumns map[int]bool, blockedUpDiagonals map[int]bool, blockedDownDiagonals map[int]bool) {
	blockedColumns[col] = true
	blockedUpDiagonals[row+col] = true
	blockedDownDiagonals[row-col] = true
}

func removeQueen(row int, col int, blockedColumns map[int]bool, blockedUpDiagonals map[int]bool, blockedDownDiagonals map[int]bool) {
	blockedColumns[col] = false
	blockedUpDiagonals[row+col] = false
	blockedDownDiagonals[row-col] = false
}
