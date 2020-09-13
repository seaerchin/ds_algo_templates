package misc

// Horner implements horner's rule for evaluating a polynomial
// horner's rule can be thought of as a right-fold in functional terms
// run-time complexity is o(n) as we iterate over exponents exactly once
func Horner(exponents []int, xVal int) int {
	y := 0
	for i := len(exponents) - 1; i >= 0; i-- {
		y = exponents[i] + xVal*y
	}
	return y
}
