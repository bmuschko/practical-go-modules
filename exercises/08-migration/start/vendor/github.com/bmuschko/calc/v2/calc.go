package calc

// Add multiple numbers.
// Return the result.
func Add(args ...int) int {
	result := 0
	for _, val := range args {
		result += val
	}
	return result
}

// Subtract two numbers.
// Return the result.
func Subtract(a, b int) int {
	return a - b
}

// Multiply two numbers.
// Return the result.
func Multiply(a, b int) int {
	return a * b
}

// Divide two numbers.
// Return the result.
func Divide(a, b int) float64 {
	return float64(a / b)
}
