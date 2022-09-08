package analyser

import (
	"fmt"
	"strconv"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/fatih/color"
)

func getSide(isCt bool) string {
	if isCt {
		return "ct"
	}
	return "t"
}

func printWithFormat(teamName string, firstHalf *Half, secondHalf *Half, index int) {
	t1side := getSide(teamName == firstHalf.ctName)
	//|t1firstside=|t1t=|t1ct=|t2t=|t2ct=
	var format []string
	if t1side == "ct" {
		format = append(format,
			color.BlueString(t1side),
			color.RedString(strconv.Itoa(secondHalf.halfTScore)),
			color.BlueString(strconv.Itoa(firstHalf.halfCtScore)),
			color.RedString(strconv.Itoa(firstHalf.halfTScore)),
			color.BlueString(strconv.Itoa(secondHalf.halfCtScore)),
		)
	} else if t1side == "t" {
		format = append(format,
			color.RedString(t1side),
			color.RedString(strconv.Itoa(firstHalf.halfTScore)),
			color.BlueString(strconv.Itoa(secondHalf.halfCtScore)),
			color.RedString(strconv.Itoa(secondHalf.halfTScore)),
			color.BlueString(strconv.Itoa(firstHalf.halfCtScore)),
		)
	}
	key := ""
	if index >= 2 {
		key = strconv.Itoa(index / 2)
		key = "o" + key
	}
	fmt.Fprintf(color.Output, "\t|"+key+"t1firstside=%s|"+
		key+"t1t=%s|"+
		key+"t1ct=%s|"+
		key+"t2t=%s|"+
		key+"t2ct=%s\n",
		format[0], format[1], format[2], format[3], format[4])
}

func (analyser *Analyser) printHalfs() {
	if len(analyser.halfs) == 0 {
		return
	}

	ctName := analyser.halfs[0].ctName
	tName := analyser.halfs[0].tName

	fmt.Printf("Team1:%s\n", ctName)
	for i := 0; i < len(analyser.halfs); i = i + 2 {
		printWithFormat(ctName, analyser.halfs[i], analyser.halfs[i+1], i)
	}

	fmt.Printf("Team1:%s\n", tName)
	for i := 0; i < len(analyser.halfs); i = i + 2 {
		printWithFormat(tName, analyser.halfs[i], analyser.halfs[i+1], i)
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

func (analyser *Analyser) printScore() {
	ctName := analyser.parser.GameState().TeamCounterTerrorists().ClanName()
	tName := analyser.parser.GameState().TeamTerrorists().ClanName()
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), analyser.ctScore, analyser.tScore)
	} else {
		fmt.Printf("%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), analyser.ctScore, analyser.tScore)
	}
}
