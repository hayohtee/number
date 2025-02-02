package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

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

// isPrime determines if a given integer is a prime number and sends the result to the provided channel.
//
// Parameters:
//   - value: The integer to be checked for.
//   - resultChan: A channel to send the boolean result indicating whether the number is prime or not.
//
// The function first handles edge cases for numbers less than or equal to 3.
// It then checks divisibility by 2 and 3.
// For larger numbers, it uses a loop to check divisibility by numbers of the form 6k ± 1 up to the square root of the value.
func isPrime(value int, resultChan chan<- bool) {
	if value <= 1 {
		resultChan <- false
		return
	}

	if value <= 3 {
		resultChan <- true
		return
	}

	if value%2 == 0 || value%3 == 0 {
		resultChan <- false
		return
	}

	sqrt := int(math.Sqrt(float64(value)))

	for i := 5; i <= sqrt; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			resultChan <- false
		}
	}
	resultChan <- true
}

// isArmstrong checks if a given integer is an Armstrong number.
// An Armstrong number (also known as a narcissistic number) is a number that is equal to the sum of its own digits each raised to the power of the number of digits.
// For example, 153 is an Armstrong number because 1^3 + 5^3 + 3^3 = 153.
//
// Parameters:
//   - value: The integer to be checked.
//   - resultChan: A channel to send the result (true if the number is an Armstrong number, false otherwise).
//
// The function sends the result to the provided resultChan channel.
func isArmstrong(value int, resultChan chan<- bool) {
	if value < 0 {
		resultChan <- false
		return
	}

	temp := value
	numDigits := 0

	for temp > 0 {
		numDigits++
		temp /= 10
	}

	sum := 0
	temp = value
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	resultChan <- sum == value
}

// getFunFact fetches a fun fact about a given number from the Numbers API.
// It takes an integer value as input and returns a string containing the fun fact
// or an error if the request fails or the context times out.
//
// The function creates a context with a timeout of 1 second to ensure the request
// does not hang indefinitely. It then constructs an HTTP GET request to the Numbers API
// using the provided number. If the request is successful, it reads the response body
// and returns it as a string. If any error occurs during the process, it returns an empty
// string and the error.
func getFunFact(value int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://numbersapi.com/%d", value), nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
