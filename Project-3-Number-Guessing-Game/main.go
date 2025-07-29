package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num int
	fmt.Println("Number Has Been Gussed - Enter Yours !")
	var guss int = rand.Intn(10) + 1

	for {

		fmt.Scan(&num)

		if num == guss {
			fmt.Println("You Have Gussed Correctly ! Yay")
			break
		} else {
			fmt.Println("Your Gussed Wrong Try Again. ;(")
		}
	}

	fmt.Println("See Ya Later")

}
