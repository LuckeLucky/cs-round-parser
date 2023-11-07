package analyser

import (
	"fmt"

	"github.com/LuckeLucky/cs-round-parser/utils"
	"github.com/fatih/color"
)

func (analyser *Analyser) printMap() {
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "Map:%s\n", color.YellowString(analyser.mapName))
	} else {
		fmt.Printf("Map:%s\n", color.YellowString(analyser.mapName))
	}
}

func printParticipant(steamID uint64, name string) {
	fmt.Printf("https://steamcommunity.com/profiles/%d (%s)\n", steamID, name)
}

func (analyser *Analyser) printSpectators() {
	fmt.Println("----------------------------")
	fmt.Println("Spectators:")
	for steamID, name := range analyser.spectators {
		printParticipant(steamID, name)
	}
}

func (analyser *Analyser) printPlayers() {
	for teamName, players := range analyser.players {
		fmt.Println("----------------------------")
		fmt.Printf("%s:\n", teamName)
		for steamID, name := range players {
			printParticipant(steamID, name)
		}
	}
}
