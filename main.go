package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	num1, num2 := getTwoNumbersFromUser(reader)
	fmt.Println(num1, num2)
	nextNum := calculateNextNumber(num1, num2)
	fmt.Printf("The next number in the sequence is: %d\n", nextNum)
	fmt.Println("The end.")
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
				if err1 == nil && err2 == nil {
					fmt.Printf("The two integers are: %d and %d\n", num1, num2)
					return num1, num2
				} else {
					fmt.Println("Error: Invalid input, please enter two integers.")
				}
			} else {
				fmt.Println("Error: Invalid input, please enter two integers.")
			}
		}
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
	for i := 2; i < num; i++ {
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
