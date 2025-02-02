package main

import "math"

// isPerfectNumber checks if a given integer is a perfect number.
// A perfect number is a positive integer that is equal to the sum of its proper divisors,
// excluding the number itself. For example, 6 is a perfect number because its proper divisors
// are 1, 2, and 3, and 1 + 2 + 3 = 6.
//
// Parameters:
//
//	value (int): The integer to check.
//
// Returns:
//
//	bool: True if the integer is a perfect number, false otherwise.
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

// isPrime checks if a given integer is a prime number.
// It returns true if the number is prime, and false otherwise.
//
// A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.
//
// The function first handles simple cases where the number is less than or equal to 3.
// For numbers greater than 3, it checks divisibility by 2 and 3.
// Then, it uses a loop to check for factors up to the square root of the number, skipping even numbers.
//
// Parameters:
//   - value: the integer to be checked for primality.
//
// Returns:
//   - bool: true if the number is prime, false otherwise.
func isPrime(value int) bool {
	if value <= 1 {
		return false
	}

	if value <= 3 {
		return true
	}

	if value%2 == 0 || value%3 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(value)))

	for i := 5; i <= sqrt; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
