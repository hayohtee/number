package main

import "math"

// isPerfect checks if a given integer is a perfect number.
// A perfect number is a positive integer that is equal to the sum of its proper divisors, excluding itself.
// For example, 6 is a perfect number because its divisors are 1, 2, and 3, and 1 + 2 + 3 = 6.
//
// Parameters:
//
//	value (int): The integer to check.
//
// Returns:
//
//	bool: True if the integer is a perfect number, false otherwise.
func isPerfect(value int) bool {
	if value <= 1 {
		return false
	}

	sum := 1

	for i := 2; i*i <= value; i++ {
		if value%i == 0 {
			sum += i
			if i != value/i {
				sum += value / i
			}
		}
	}

	return sum == value
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
