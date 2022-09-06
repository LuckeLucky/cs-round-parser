package analyser

import (
	"fmt"
	"strconv"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/fatih/color"
)

func (analyser *Analyser) printHalfs() {
	if len(analyser.halfs) == 0 {
		return
	}

	var firstCtNme string

	for i, half := range analyser.halfs {
		if i == 0 {
			firstCtNme = half.ctName
			fmt.Printf("Half  |%.6s|%.6s\n", utils.PadSpaceEnd(half.ctName, 6), utils.PadSpaceEnd(half.tName, 6))
		}
		printScoresHalf(half, i+1, half.ctName == firstCtNme)
	}
}

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

func printScoresHalf(half *Half, nHalf int, isCTOnLeft bool) {
	ctScore := utils.PadSpaceEnd(strconv.Itoa(half.halfCtScore), 6)
	tScore := utils.PadSpaceEnd(strconv.Itoa(half.halfTScore), 6)

	halfPrint := utils.PadSpaceEnd(strconv.Itoa(nHalf), 6)
	format := []string{color.RedString(tScore), color.BlueString(ctScore)}
	if isCTOnLeft {
		format[0], format[1] = format[1], format[0]
	}
	fmt.Printf("%s|", halfPrint)
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "%s|%s\n", format[0], format[1])
	} else {
		fmt.Printf("%s|%s\n", format[0], format[1])
	}
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

func printParticipant(steamID uint64, name string) {
	fmt.Printf("https://steamcommunity.com/profiles/%d (%s)\n", steamID, name)
}

func (analyser *Analyser) printSpectators() {
	fmt.Println("Spectators:")
	for steamID, name := range analyser.spectators {
		printParticipant(steamID, name)
	}
}

func (analyser *Analyser) printPlayers() {
	fmt.Println("Players:")
	for teamName, players := range analyser.players {
		fmt.Printf("%s:\n", teamName)
		for steamID, name := range players {
			printParticipant(steamID, name)
		}
	}
}
