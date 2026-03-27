package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	//var help string
	var operator string

	for {

		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)

		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		fmt.Print("operator :")
		fmt.Scan(&operator)
		if operator == "*, /, -, +" {
			fmt.Println("valid operator")
		}
		if operator == "quit" {
			fmt.Println("exit!")
			break
		}
		if operator == "help" {
			fmt.Println("your input should be:")
			fmt.Println("* = multiply")
			fmt.Println("/ = div")
			fmt.Println("- = subtract")
			fmt.Println("+ = add")
		}

		switch operator {
		case "*":
			fmt.Println("Multiply :", num1, num2)
			fmt.Println("calculate :", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Can't divide by zero")
				continue
			}
			fmt.Println("div :", num1, num2)
			fmt.Println("Result :", num1/num2)
		case "-":
			fmt.Println("sub :", num1, num2)
			fmt.Println("Result :", num1-num2)
		case "+":
			fmt.Println("Add :", num1, num2)
			fmt.Println("Result :", num1+num2)
		case "help":

		default:
			fmt.Println("No valid operation entered")
		}
		continue

	}

}
