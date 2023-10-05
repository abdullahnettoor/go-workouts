//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestHealth(t *testing.T) {

	player1 := Player{name: "test1", health: 95, maxHealth: 100}

	player1.increaseHealth(25)
	if player1.health > player1.maxHealth {
		t.Fatalf("Health went over limit: %v want: %v", player1.health, player1.maxHealth)
	}
	player1.decreaseHealth(player1.maxHealth + 6)
	if player1.health < 0 {
		t.Fatalf("Health: %v Minimum: 0", player1.health)
	}
	if player1.health > player1.maxHealth {
		t.Fatalf("Health: %v Maximum: %v", player1.health, player1.maxHealth)
	}

}
