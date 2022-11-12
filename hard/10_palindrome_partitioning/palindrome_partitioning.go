package main

import (
	"fmt"
)

func palindromePartitioningMinCuts(inputString string) {
	palindromes := make([][]bool, len(inputString))
	for i := 0; i < len(inputString); i++ {
	  row := make([]bool, len(inputString))
	  for j := 0; j < len(inputString); j++ {
		if i == j {
		  row[j] = true
		} else {
		  row[j] = false
		}
	  }
	  palindromes[i] = row
	}
	for length := 2; length < len(inputString)+1; length++ {
	  for i := 0; i < len(inputString)-length+1; i++ {
		j := i + length - 1
		if length == 2 {
		  palindromes[i][j] = inputString[i] == inputString[j]
		} else {
		  palindromes[i][j] = inputString[i] == inputString[j] && palindromes[i+1][j-1]
		}
	  }
	}
	cuts := make([]int, len(inputString))
	for i := 0; i < len(inputString); i++ {
	  cuts[i] = math.MaxInt64
	}
	for i := 0; i < len(inputString); i++ {
	  if palindromes[0][i] {
		cuts[i] = 0
	  } else {
		cuts[i] = cuts[i-1] + 1
		for j := 1; j < i; j++ {
		  if palindromes[j][i] && cuts[j-1]+1 < cuts[i] {
			cuts[i] = cuts[j-1] + 1
		  }
		}
	  }
	}
	return cuts[len(cuts)-1]
  }

  fmt.Println(palindromePartitioningMinCuts('noonabbad'))
  fmt.Println(palindromePartitioningMinCuts('ababbbabbababa'))
