package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"strconv"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	num1, num2 := getTwoNumbersFromUser(reader)
	fmt.Println(num1, num2)
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
