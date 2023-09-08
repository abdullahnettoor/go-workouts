//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

const (
	SmallLift = iota - 1
	StandardLift
	LargeLift
)

type Lift int

type Motorcycle string
type Car string
type Truck string
type LiftPicker interface {
	PickLift() Lift
}

// func (m Motorcycle) Name() string {
// 	return fmt.Sprintf("Motorcycle : %v", string(m))
// }
// func (m Car) Name() string {
// 	return fmt.Sprintf("Car : %v", string(m))
// }
// func (m Truck) Name() string {
// 	return fmt.Sprintf("Truck : %v", string(m))
// }

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}
func (m Car) PickLift() Lift {
	return StandardLift
}
func (m Truck) PickLift() Lift {
	return LargeLift
}

func (c Car) ShowTopSpeed(speed int) {
	fmt.Println("The Speed of", c, "is", speed)
}

func sendToLift(l LiftPicker) {
	switch l.PickLift() {
	case SmallLift:
		fmt.Println("Send to Small Lift, It's", l)
	case StandardLift:
		fmt.Println("Send to Standard Lift, It's", l)
	case LargeLift:
		fmt.Println("Send to Large Lift, It's", l)
	}
}

func main() {
	car := Car("Mercedez")
	truck := Truck("Bharath Benz")
	bike := Motorcycle("Ninja")

	car.ShowTopSpeed(120)

	sendToLift(car)
	sendToLift(truck)
	sendToLift(bike)
}
