package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"log"
	"os"
)

func main() {
	trelloKey := os.Getenv("TRELLO_KEY")
	trelloToken := os.Getenv("TRELLO_TOKEN")
	trelloUser := os.Getenv("TRELLO_USER")

	trello, err := trello.NewAuthClient(trelloKey, &trelloToken)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member(trelloUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trello Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}
