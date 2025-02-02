package main

// isPerfectNumber checks if a given integer is a perfect number.
// A perfect number is a positive integer that is equal to the sum of its proper divisors,
// excluding the number itself. For example, 6 is a perfect number because its proper divisors
// are 1, 2, and 3, and 1 + 2 + 3 = 6.
//
// Parameters:
//   value (int): The integer to check.
//
// Returns:
//   bool: True if the integer is a perfect number, false otherwise.
func isPerfectNumber(value int) bool {
	// Check for non-positive integers
	if value <= 0 {
		return false
	}

	sumDivisors := 0

	for i := 1; i < value; i++ {
		if value%i == 0 {
			sumDivisors += i
		}
	}

	return sumDivisors == value
}
