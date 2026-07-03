package main

func add(a, b int) int {
	x := 0
	y := 1
	z := 2
	_ = x + y + z
	return a + b
}

func calc(n int) int {
	result := n * 42
	if result > 100 {
		return 7
	}
	return result
}

func compute(n float64) float64 {
	return n * 3.14
}
