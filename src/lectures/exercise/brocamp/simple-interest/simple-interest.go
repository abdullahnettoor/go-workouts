// 3. Write a program to find the simple interest.
// Program should accept 3 inputs from the user and calculate simple interest for the given inputs. Formula: SI=(P*R*n)/100)
// Variable
// Data Type
// Principal amount (P)
// Integer
// Interest rate (R)
// Float
// Number of years (n)
// Float
// Simple Interest (SI)
// Float

package main

import "fmt"

type PrincipalAmt int
type InterestRate float32
type NoOfyears float32

func main() {
	var (
		PrincipalAmt   int
		InterestRate   float32
		NoOfyears      float32
		SimpleInterest float32
	)

	fmt.Println("Enter details to find Simple Interest")
	fmt.Printf("Principal Amount :\t")
	fmt.Scanf("%d", &PrincipalAmt)
	fmt.Printf("Rate of Interest :\t")
	fmt.Scanf("%f", &InterestRate)
	fmt.Printf("No. of Years :\t")
	fmt.Scanf("%f", &NoOfyears)

	SimpleInterest = (float32(PrincipalAmt) * InterestRate * NoOfyears) / 100

	fmt.Println("Simple interest is", SimpleInterest)
}
