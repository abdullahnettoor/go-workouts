//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import (
	"fmt"
)

const (
	active   = true
	inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activateTag(tag *SecurityTag) {
	*tag = active
}

func deactivateTag(tag *SecurityTag) {
	*tag = inactive
}

func checkout(list []Item) {
	for i, _ := range list {
		deactivateTag(&list[i].tag)
	}
}

func printItemsList(title string, list []Item) {
	fmt.Println(title)
	for _, item := range list {
		fmt.Println(item.name, "\t:", item.tag)
	}
	fmt.Println()
}

func main() {

	shirt := Item{"Shirt", active}
	pants := Item{"Pants", active}
	cap := Item{"Cap", active}
	shoes := Item{"Shoes", active}

	itemList := []Item{shirt, pants, cap}
	itemList = append(itemList, shoes)

	printItemsList("Items Added", itemList)

	deactivateTag(&itemList[2].tag)

	printItemsList("Cart Updated", itemList)

	checkout(itemList)
	printItemsList("Cheking Out ...", itemList)

}
