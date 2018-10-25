package main

import "fmt"

func main() {
	fnMax := max(1, 3)
	fmt.Println(fnMax())

	defer func() {
		fmt.Println("End execution")
		str := recover()
		fmt.Println(str)
	}()

	testFunction([]float64{1, 3, 5})
}

func testFunction(xs []float64) float64 {
	panic("Not Implemented")
}

func max(a int, b int) (r func() int) {
	fn := func() int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	r = fn

	return
}
