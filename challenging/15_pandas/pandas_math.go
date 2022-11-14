package main

import (
	"fmt"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

func main() {
	ds1 := mat64.NewVector(5, []float64{2, 4, 6, 8, 10})
	ds2 := mat64.NewVector(5, []float64{1, 3, 5, 7, 9})

	add := mat64.NewVector(5, make([]float64, 5))
	add.AddVec(ds1, ds2)

	sub := mat64.NewVector(5, make([]float64, 5))
	sub.SubVec(ds1, ds2)

	mul := mat64.NewVector(5, make([]float64, 5))
	floats.MulTo(mul.RawVector().Data, ds1.RawVector().Data, ds2.RawVector().Data)

	div := mat64.NewVector(5, make([]float64, 5))
	floats.DivTo(div.RawVector().Data, ds1.RawVector().Data, ds2.RawVector().Data)

	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
