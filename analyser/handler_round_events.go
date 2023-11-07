package analyser

import (
	"fmt"

	"github.com/LuckeLucky/cs-round-parser/utils"
	"github.com/fatih/color"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

func (analyser *Analyser) registerMatchEventHandlers() {
	analyser.parser.RegisterEventHandler(func(e events.RoundEnd) { analyser.handleRoundEnd(e) })
}

func (analyser *Analyser) handleRoundEnd(e events.RoundEnd) {
	_, err := analyser.getGameTick()
	if err {
		return
	}

	//Score not updated in source
	winnerScore := e.WinnerState.Score()
	loserScore := e.LoserState.Score()
	if !analyser.isSource2 {
		winnerScore = winnerScore + 1
	}
	switch e.Winner {
	case common.TeamCounterTerrorists:
		analyser.printScore(winnerScore, loserScore)
	case common.TeamTerrorists:
		analyser.printScore(loserScore, winnerScore)
	}
	analyser.setParticipants()
}

func (analyser *Analyser) printScore(ctScore int, tScore int) {
	ctName := analyser.getTeamName(analyser.parser.GameState().TeamCounterTerrorists())
	tName := analyser.getTeamName(analyser.parser.GameState().TeamTerrorists())
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), ctScore, tScore)
	} else {
		fmt.Printf("%s vs %s  %d : %d\n", color.BlueString(ctName), color.RedString(tName), ctScore, tScore)
	}
}
