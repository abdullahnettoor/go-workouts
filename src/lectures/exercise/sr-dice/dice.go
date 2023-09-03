//  --Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//  --Requirements:
//  * Print the sum of the dice roll
//  * Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//  * The program must handle any number of dice, rolls, and sides
//
//  --Notes:
//  * Use packages from the standard library to complete the project

package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func diceRoll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 12
	rolls := 2

	Println("There is", dice, "dices and each has", sides, "sides")
	Println("We are going to roll it", rolls, "times")
	Println()

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := diceRoll(sides)
			sum += rolled
			Println("Roll #", r, "of Dice", d, "gives", rolled)
		}
		if sum == 2 && dice == 2 {
			//  - "Snake eyes": when the total roll is 2, and total dice is 2
			Println("Snake Eyes")
		} else if sum == 7 {
			//  - "Lucky 7": when the total roll is 7
			Println("Lucky 7")
		} else if sum%2 == 0 {
			//  - "Even": when the total roll is even
			Println("Even")
		} else if sum%2 == 1 {
			//  - "Odd": when the total roll is odd
			Println("Odd")
		}
		Println("Sum of Roll", r, "is", sum)
		Println()
	}
}
