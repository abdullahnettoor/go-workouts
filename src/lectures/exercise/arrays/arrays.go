//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func showShoppingList(list [4]Product) {
	fmt.Println()

	var cost, totalItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			totalItems++
			cost += item.price
		}
	}
	fmt.Println("Last Item is", list[totalItems-1])
	fmt.Println("Shopping List contains", totalItems, "items.")
	fmt.Println("Total cost will be", cost, ".")
	fmt.Println()
}

func main() {

	var shoppingList = [4]Product{}

	shoppingList[0] = Product{"iPhone 15 Pro Ultra", 1200}
	shoppingList[1] = Product{"iPad 10", 800}
	shoppingList[2] = Product{"MacBook Pro", 1800}

	showShoppingList(shoppingList)

	shoppingList[3] = Product{"Apple Watch", 1000}
	showShoppingList(shoppingList)

}
