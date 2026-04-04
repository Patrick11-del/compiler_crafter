package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var choice int
		var err error
		var input string

		fmt.Print()
		fmt.Println("Base Converter: ")
		for {
			fmt.Println("====Menue Section====")

			flin := []string{
				"hexadecimal to all",
				"decimal to all",
				"binary to all",
				"exit converter",
			}
			for i, x := range flin {
				fmt.Println(i+1, x)
			}
			fmt.Print("\nselect option:  ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			choice, err = strconv.Atoi(input)
			if err != nil || choice < 1 || choice > 4 {
				fmt.Println("\nInvalid Option")
				continue
			}
			if choice == 4 {
				fmt.Println("Start Operation!")
				return
			}
			break
		}
		
		if choice == 1 {
			for {
				fmt.Println("\n>>>>Hexadecimal conversion<<<<<")
				fmt.Print("Enter Your Hex Value: ")
				hexInput, _ := reader.ReadString('\n')
				hexInput = strings.TrimSpace(hexInput)

				result, err := strconv.ParseInt(hexInput, 16, 64)
				if err != nil {
					fmt.Printf("Invalid %s Hex input\n", hexInput)
					continue
				}
				fmt.Print("Enter Base Value: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				targetBase, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Error....Enter a valid number")
					continue
				}
				if targetBase != 2 && targetBase != 10 && targetBase != 16 && targetBase != 8 {
					fmt.Println("Invalid input....can only support base 2 and 10")
					continue

				}
				resultStr := strconv.FormatInt(result, targetBase)
				fmt.Println("\n====RESULT====")
				fmt.Printf("HEX VALUE:  %s\n", hexInput)
				fmt.Printf("BASE RESULT:  %d ===> %s\n", targetBase, resultStr)
				
			}
		}
		if choice == 3 {
			for {
				fmt.Println("------ BINARY CONVERSION ------")

				fmt.Println("Enter Binary value: ")
				binInput, _ := reader.ReadString('\n')
				binInput = strings.TrimSpace(binInput)

				result, err := strconv.ParseInt(binInput, 2, 64)
				if err != nil {
					fmt.Println("ERROR ---> Invalid binary number")
					continue
				}

				fmt.Println("Enter base value: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				convertTo, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("\n[INPUT ERROR] ---->> Enter a valid base number\n")
					continue
				}

				if convertTo != 10 && convertTo != 16 && convertTo != 2 && convertTo != 8  {
					fmt.Printf("\n[ERROR] --->> '%d' is not a known, valid or existing base, Try(10, 16, 2, 8)\n", convertTo)
					continue
				}

				resultStr := strconv.FormatInt(result, convertTo)

				
				fmt.Printf("  BINARY VALUE:  '%s'\n", binInput,)
				fmt.Printf("  BASE RESULT:'%d': %s\n",  convertTo, resultStr)
				fmt.Print("============================\n")

				break

			}
		}


		if choice == 2 {
			for {
				fmt.Println("<<<<<<<DECIMAL CONVERSION>>>>>>")

				fmt.Print("Enter Deimal Value: ")
				decInput, _ := reader.ReadString('\n')
				decInput = strings.TrimSpace(decInput)

				result, err := strconv.ParseInt(decInput, 10, 64)
				if err != nil {
					fmt.Println("Input Error!---> Not a decimal value")
					continue
				}
				fmt.Print("Enter Base Number: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)

				n, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Error!---> Enter a Valid Base Number")
					continue
				}
				if n != 2 && n != 16 && n != 10 && n != 8 {
					fmt.Println("Invalid Input...can only support base 2 and base 10!")
					continue
				}
				resultStr := strconv.FormatInt(result, n)
				fmt.Println("You converted,", result, "of Base", n, "to", resultStr)

			}
		}

	}
}
