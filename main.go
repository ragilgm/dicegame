package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// define reader stdin
	reader := bufio.NewReader(os.Stdin)

	// define player
	fmt.Print("Define Player => ")
	definePlayer, _, err := reader.ReadLine()
	var players string
	for i := 0; i < len(definePlayer); i++ {
		players += string(definePlayer[i])
	}
	if err != nil {
		fmt.Println(err)
	}

	// define dice
	fmt.Print("Define Dice => ")
	var dices string
	defineDice, _, err := reader.ReadLine()
	for i := 0; i < len(defineDice); i++ {
		dices += string(defineDice[i])
	}
	if err != nil {
		fmt.Println(err)
	}

	player, err := strconv.Atoi(players)
	dice, err := strconv.Atoi(dices)

	// play game
	palayGame(reader, player, dice)
}

// object player
type Player struct {
	PlayerName  string
	PlayerPoint int
	Chance      int
	operDice    []int
	Dices       []int
	isFinish    bool
}

func palayGame(reader *bufio.Reader, definePlayer int, defineDice int) {
	var players []*Player

	// define player
	for i := 1; i <= definePlayer; i++ {
		players = append(players, &Player{PlayerName: fmt.Sprintf("Player (%s)", strconv.Itoa(i)), PlayerPoint: 0, Chance: defineDice})
	}

	// check game is end
	isEnd := false

	// begin
	for !isEnd {

		// press enter for next game
		fmt.Print("Press Enter for Next Game => ")
		reader.ReadLine()

		// generate dice
		generateDice(players)

		// before evaluate
		fmt.Println("before evaluate")
		showScore(players)

		// evaluate player
		evaluatePlayer(players)

		// after evaluate player
		fmt.Println("after evaluate")
		showScore(players)

		// chek winner player
		if checkFinish(players) {
			winner := chooseWinner(players)
			fmt.Println("The Winner is : ", winner.PlayerName)
			fmt.Println("Total Point : ", winner.PlayerPoint)
			isEnd = true
		}

	}

}

// choose winner
func chooseWinner(players []*Player) *Player {
	winnerPlayer := players[0]
	for _, player := range players {
		if player.PlayerPoint >= winnerPlayer.PlayerPoint {

			// if total point have same then choose dice equal 0 for define winner
			if player.PlayerPoint == winnerPlayer.PlayerPoint {
				if len(player.Dices) == 0 {
					winnerPlayer = player
				}
			} else {
				winnerPlayer = player
			}
		}
	}
	return winnerPlayer
}

func generateDice(players []*Player) {
	for _, player := range players {
		var dice []int
		for i := 0; i < player.Chance; i++ {
			randomNumber := rand.Intn(6) + 1
			dice = append(dice, randomNumber)
		}
		player.Dices = dice
	}
}

func evaluatePlayer(players []*Player) {
	for idx, player := range players {
		var newDice []int
		for i := 0; i < len(player.Dices); i++ {
			dice := player.Dices[i]
			if dice == 6 {
				player.PlayerPoint += 1
				continue
			} else if dice == 1 {

				if len(players)-1 == idx {
					players[0].operDice = append(players[0].operDice, dice)
				} else {
					players[idx+1].operDice = append(players[idx+1].operDice, dice)
				}

			} else {
				newDice = append(newDice, dice)
			}
		}
		player.Dices = newDice
		player.Chance = len(newDice)
	}

	for _, player := range players {
		for i := 0; i < len(player.operDice); i++ {
			player.Dices = append(player.Dices, player.operDice[i])
		}
		player.operDice = []int{}
		player.Chance = len(player.Dices)
		if player.Chance == 0 {
			player.isFinish = true
		}
	}

}

func checkFinish(players []*Player) bool {
	countPlayer := 0
	for _, player := range players {
		if !player.isFinish {
			countPlayer += 1
		}
	}
	if countPlayer == 1 {
		return true
	}
	return false
}

func showScore(players []*Player) {
	fmt.Println("===================================")
	for _, player := range players {
		fmt.Println("Player Name : " + player.PlayerName)
		fmt.Println("Point : " + strconv.Itoa(player.PlayerPoint))
		fmt.Println("Chance : " + strconv.Itoa(player.Chance))
		fmt.Println("Dice : ", player.Dices)
	}
	fmt.Println("")
}
