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

func (analyser *Analyser) printFinish() {
	fmt.Println("---Finish---")
}

func (analyser *Analyser) printHalf() {
	fmt.Println("---HALF---")
}

func (analyser *Analyser) printRoundsPlayed() {
	fmt.Printf("Rounds played:%d\n", analyser.roundsPlayed)
}

func (analyser *Analyser) printScore() {
	ctName := analyser.parser.GameState().TeamCounterTerrorists().ClanName()
	tName := analyser.parser.GameState().TeamTerrorists().ClanName()
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), analyser.ctScore, analyser.tScore)
	} else {
		fmt.Printf("%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), analyser.ctScore, analyser.tScore)
	}
}
