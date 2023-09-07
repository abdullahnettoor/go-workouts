//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
}

func (player *Player) increaseHealth(health uint) {
	if player.health == 100 {
		fmt.Println(player.name, "already has max health -> Health :", player.health)
	} else if player.health+health > player.maxHealth {
		player.health = player.maxHealth
		fmt.Println(player.name, "gained max health -> Health :", player.health)
	} else {
		player.health += health
		fmt.Println(player.name, "gained", health, "health -> Health :", player.health)
	}
}

func (player *Player) decreaseHealth(health uint) {
	if health <= player.health {
		player.health -= health
		fmt.Println(player.name, "lost", health, "health -> Health :", player.health)

	} else {
		player.health = 0
		fmt.Println(player.name, "dead -> Health :", player.health)
	}
}

func main() {
	player1 := Player{
		name:      "Savio",
		health:    100,
		maxHealth: 100,
	}

	player1.decreaseHealth(25)
	player1.increaseHealth(20)
	player1.increaseHealth(30)
	player1.decreaseHealth(85)
	player1.decreaseHealth(20)
}
