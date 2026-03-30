package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the CLI Calculator!")
	fmt.Println("Type 'help' to start.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "help":
			printHelp()

		case "quit":
			fmt.Println("Goodbye!")
			return

		case "add", "sub", "mul", "div", "mod":
			handleOperation(command, parts)

		default:
			fmt.Println("Unknown command. Type 'help'.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <a> <b>   → addition")
	fmt.Println("  sub <a> <b>   → subtraction")
	fmt.Println("  mul <a> <b>   → multiplication")
	fmt.Println("  div <a> <b>   → division")
	fmt.Println("  mod <a> <b>   → module")
	fmt.Println("  help          → show this message")
	fmt.Println("  quit          → exit the program")
}

func handleOperation(command string, parts []string) {
	if len(parts) != 3 {
		fmt.Println("Wrong number of arguments. Format: <op> <a> <b>")
		return
	}

	a, err1 := strconv.ParseFloat(parts[1], 64)
	b, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Both arguments must be numbers.")
		return
	}

	var result float64

	switch command {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "mod":
		result = math.Mod(a, b)
	case "pow":
		result = math.Pow(a, b)
	case "div":
		if b == 0 {

			fmt.Println("Division by zero is not allowed.")
			return
		}
		result = a / b
	}

	fmt.Printf("✦ Result: %.2f\n", result)
}
