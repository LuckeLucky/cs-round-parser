package analyser

import (
	"github.com/LuckeLucky/cs-round-parser/global"
	"github.com/LuckeLucky/cs-round-parser/utils"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

const (
	MAX_ROUNDS_REGULAR = 30
	WIN_ROUNDS_REGULAR = MAX_ROUNDS_REGULAR/2 + 1
)

func (analyser *Analyser) checkValidRoundStartMoney() bool {
	// if the money value is not set, no need to check
	if !analyser.isMoneySet {
		return true
	}

	// between 0 - 30 rounds start money is 800
	if analyser.roundsPlayed < 30 {
		return analyser.currentStartMoney == global.RegularStartMoney
	} else {
		// when overtime money isnt set we can't say if there is a valid ot start moneys
		if !analyser.isOvertimeMoneySet {
			return true
		}
		return analyser.currentOvertimeStartMoney == global.OvertimeStartMoney
	}

}

func (analyser *Analyser) checkMatchHalf() bool {
	if analyser.roundsPlayed == MAX_ROUNDS_REGULAR/2 {
		return true
	}

	ctScore, tScore := analyser.ctScore, analyser.tScore
	// overtime
	roundsInOvertime := ctScore + tScore - MAX_ROUNDS_REGULAR
	if roundsInOvertime == 0 && ctScore == tScore {
		return true
	} else if roundsInOvertime > 0 {
		return roundsInOvertime%(analyser.overtimeMaxRounds/2) == 0
	}
	return false
}

func (analyser *Analyser) checkMatchFinished() bool {
	ctScore, tScore := analyser.ctScore, analyser.tScore
	roundsInOvertime := ctScore + tScore - MAX_ROUNDS_REGULAR

	if ((ctScore == WIN_ROUNDS_REGULAR) != (tScore == WIN_ROUNDS_REGULAR)) || roundsInOvertime >= 0 {
		absDiff := utils.Abs(ctScore - tScore)
		x := roundsInOvertime % analyser.overtimeMaxRounds
		nRoundsOfOTHalf := analyser.overtimeMaxRounds / 2
		if roundsInOvertime < 0 || ((x == 0 && absDiff == 2) || (x > nRoundsOfOTHalf && absDiff >= nRoundsOfOTHalf)) {
			return true
		}
	}
	return false
}

func (analyser *Analyser) checkMatchEnded() bool {
	return analyser.matchEnded
}

func (analyser *Analyser) checkFreeArmor() bool {
	return analyser.freeArmor == 0
}

func (analyser *Analyser) isPreGame() bool {
	return analyser.parser.GameState().GamePhase() == common.GamePhasePregame
}

func (analyser *Analyser) isScoreEmpty() bool {
	//T and CT start with 1k money in first Round
	if analyser.roundsPlayed > 0 {
		return false
	}

	for _, participant := range analyser.parser.GameState().Participants().AllByUserID() {
		if participant.Team > 1 && participant.Entity != nil {
			sumScore := participant.Kills() + participant.Deaths() + participant.Assists() + participant.Score()
			if sumScore > 0 {
				return false
			}

		}
	}
	return true
}

func (analyser *Analyser) checkForMatchHalfOrEnd() {
	isEnd, isHalf := analyser.checkMatchFinished(), analyser.checkMatchHalf()
	if isEnd || isHalf {
		analyser.setNewHalf()
		if isEnd {
			analyser.setMatchEnded()
			analyser.printFinish()
		} else {
			analyser.resetHalfScores()
			analyser.printHalf()
		}
	}
}
