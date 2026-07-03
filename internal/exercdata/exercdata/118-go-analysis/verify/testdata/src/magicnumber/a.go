package main

func add(a, b int) int {
	// These are NOT magic numbers (0, 1, 2 are allowed)
	x := 0
	y := 1
	z := 2
	_ = x + y + z
	return a + b
}

func calc(n int) int {
	// Magic number: 42, 100, 7
	result := n * 42
	if result > 100 {
		return 7
	}
	return result
}

func compute(n float64) float64 {
	// All float literals are magic numbers in our analyzer
	return n * 3.14
}
