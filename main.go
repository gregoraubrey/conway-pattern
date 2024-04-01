package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	num1, num2 := getTwoNumbersFromUser(reader)
	reader2 := bufio.NewReader(os.Stdin)
	length := getLengthFromUser(reader2)
	sequence := buildSequence(num1, num2, length)
	fmt.Println(sequence)
}

func getTwoNumbersFromUser(reader io.Reader) (int, int) {
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Enter two positive integers separated by a space: ")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			parts := strings.Split(input, " ")

			if len(parts) == 2 {
				num1, err1 := strconv.Atoi(parts[0])
				num2, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil && num1 > 0 && num2 > 0 {
					return num1, num2
				}
			}
		}
		fmt.Println("Error: Invalid input, please enter two integers.")
	}
}

func calculateNextNumber(num1 int, num2 int) int {
	// Calculate whether num1 and num2's sum is prime
	if isPrime(num1 + num2) {
		return num1 + num2
	}
	// Otherwise, extract the lowest prime factor of num1 + num2
	lpf := getLowestPrimeFactor(num1 + num2)
	// Then return the sum of num1 and num2 divided by its lowest prime factor
	return (num1 + num2) / lpf
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	squareRoot := int(math.Sqrt(float64(num)))
	for i := 3; i <= squareRoot; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getLowestPrimeFactor(num int) int {
	for i := 2; i < num; i++ {
		if num%i == 0 && isPrime(i) {
			return i
		}
	}
	fmt.Println("Error: No prime factor found. Returning 1 instead.")
	return 1
}
func buildSequence(num1 int, num2 int, length int) []int {
	sequence := []int{num1, num2}
	nextNum := calculateNextNumber(num1, num2)
	// Continue adding to the sequence until it contains `length` elements
	for len(sequence) < length {
		sequence = append(sequence, nextNum)
		// Call calculateNextNumber with the last two elements of the sequence
		nextNum = calculateNextNumber(sequence[len(sequence)-2], sequence[len(sequence)-1])
	}
	return sequence
}

func getLengthFromUser(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Enter the length of the sequence as a positive integer of at least 2: ")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			length, err := strconv.Atoi(input)
			if err == nil && length >= 2 {
				return length
			}
		}
		fmt.Println("Error: Invalid input, please enter a positive integer of at least 2.")
	}
}
