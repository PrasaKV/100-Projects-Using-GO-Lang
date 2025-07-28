package main

import "fmt"

func main() {

	var num1 float32
	var num2 float32
	var operator string

	for range 5 {

		fmt.Println("\nThis is A Basic Calculator")
		fmt.Println("Enter Your First Number - ")
		fmt.Scanln(&num1)
		fmt.Println("Enter Operation - ")
		fmt.Scanln(&operator)
		fmt.Println("Enter Your Second Number - ")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Printf("Anwser is - %v", add(num1, num2))
		case "-":
			fmt.Printf("Anwser is - %v", substract(num1, num2))
		case "*":
			fmt.Printf("Anwser is - %v", multiply(num1, num2))
		case "/":
			fmt.Printf("Anwser is - %v", divide(num1, num2))
		default:
			fmt.Println("Somthing is Wrong")
		}

	}

}

func add(num1, num2 float32) float32 {
	return num1 + num2
}
func substract(num1, num2 float32) float32 {
	return num1 - num2
}
func multiply(num1, num2 float32) float32 {
	return num1 * num2
}
func divide(num1, num2 float32) float32 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}
