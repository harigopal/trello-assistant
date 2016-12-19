package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	trelloKey := os.Getenv("TRELLO_KEY")
	trelloToken := os.Getenv("TRELLO_TOKEN")
	trelloBoardId := os.Getenv("TRELLO_BOARD_ID")

	if trelloKey == "" || trelloToken == "" || trelloBoardId == "" {
		log.Fatal("TRELLO_KEY, TRELLO_TOKEN, and TRELLO_BOARD_ID must be set.")
		os.Exit(1)
	}

	// Create Trello client.
	trello, err := trello.NewAuthClient(trelloKey, &trelloToken)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Load the required board.
	fmt.Print("Loading board... ")
	board, err := trello.Board(trelloBoardId)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("Done. Loaded", board.Name)
	}

	// Load all lists on the board.
	fmt.Print("Loading all lists... ")
	lists, err := board.Lists()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("Done. Loaded", len(lists), "lists.")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"List", "Number of Cards", "Total Score", "Easy (1)", "Average (2)", "Hard (3)"})

	var totalCards int

	fmt.Print("Loading cards stats")
	for _, list := range lists {
		row := make([]string, 6)
		var easyCount, averageCount, hardCount int

		// Load all cards in list.
		cards, _ := list.Cards()
		fmt.Print(".")
		totalCards += len(cards)

		for _, card := range cards {
			var score byte
			strippedName := strings.TrimSpace(card.Name)

			nameLength := len(strippedName)

			if strippedName[nameLength-1] == ')' {
				score = strippedName[nameLength-2]
				if score == '+' {
					score = strippedName[nameLength-3]
				}

				if score == '1' {
					easyCount++
				} else if score == '2' {
					averageCount++
				} else if score == '3' {
					hardCount++
				}
			}
		}

		row[0] = list.Name
		row[1] = strconv.Itoa(len(cards))
		row[2] = strconv.Itoa(easyCount + averageCount*2 + hardCount*3)
		row[3] = strconv.Itoa(easyCount)
		row[4] = strconv.Itoa(averageCount)
		row[5] = strconv.Itoa(hardCount)
		table.Append(row)
	}
	fmt.Println(" Done. Evaluated", totalCards, "cards.")

	table.Render()
}
