// Golang Basic Lesson 13-2
// Golang Basic Lesson 13-2
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	cardNumbers   int = 10
	attackNumbers int = 6
)

var (
	gamer1Cards [10]int = [10]int{4, 5, 0, 2, 8, 9, 7, 6, 3, 1}
	gamer2Cards [10]int = [10]int{6, 2, 9, 3, 1, 4, 5, 0, 7, 8}
)

func main() {
	rand.Seed(time.Now().Unix())
	ayrac := strings.Repeat("=", 30)

	fmt.Println("Game starting!!")
	fmt.Println(ayrac)
	gamer1Attack := 0
	gamer2Attack := 0

	fmt.Println("Gamer 1 card selection:")
	fmt.Scanf("%d", &gamer1Attack)
	fmt.Println("Gamer 2 card selection:")
	fmt.Scanf("%d", &gamer2Attack)

	fmt.Printf("Gamer 1 thrown dice: %d\n", gamer1Attack)
	fmt.Printf("Gamer 2 thrown dice: %d\n", gamer2Attack)
	fmt.Println(ayrac)

	gameWinningPlayer := 0
	for i := 0; i < attackNumbers; i++ {
		time.Sleep(time.Second)
		gamer1Attack = gamer2Cards[gamer1Attack]
		gamer2Attack = gamer1Cards[gamer2Attack]
		fmt.Printf("The card chosen by Player 1: %d\n", gamer1Attack)
		fmt.Printf("The card chosen by Player 2: %d\n", gamer2Attack)
		fmt.Println(ayrac)

		if gamer1Attack == 0 && gamer2Attack == 0 {
			gameWinningPlayer = 0
			break
		}

		if gamer1Attack == 0 {
			gameWinningPlayer = 1
			break
		}

		if gamer2Attack == 0 {
			gameWinningPlayer = 2
			break
		}
	}
	if gameWinningPlayer == 0 {

		if gamer1Attack < gamer2Attack {
			gameWinningPlayer = 1
		} else if gamer2Attack < gamer1Attack {
			gameWinningPlayer = 2
		} else {
			gameWinningPlayer = 0
		}
	}
	fmt.Println(ayrac)
	fmt.Println("Game finish!")
	if gameWinningPlayer == 0 {
		fmt.Println("Game is tied!")
	} else {
		fmt.Printf("Player %d. won the game\n", gameWinningPlayer)
	}
}
