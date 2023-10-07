package analyser

func (analyser *Analyser) setDefault() {
	analyser.maxRounds = 30
	analyser.overtimeMaxRounds = 6
	analyser.freeArmor = 0
	analyser.matchEnded = false
	analyser.spectators = make(map[uint64]string)
	analyser.players = make(map[string]map[uint64]string)
}

func (analyser *Analyser) resetHalfScores() {
	analyser.halfCtScore = 0
	analyser.halfTScore = 0
}

func (analyser *Analyser) switchSideScores() {
	analyser.ctScore, analyser.tScore = analyser.tScore, analyser.ctScore
}
