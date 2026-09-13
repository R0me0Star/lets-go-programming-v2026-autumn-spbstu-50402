package main

import "fmt"

func main() {
	var a, b int
	var op string

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("First operand is invalid")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Second operand is invalid")
		return
	}

	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Operation is invalid")
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero is prohibited")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Operation is invalid")
	}
}
