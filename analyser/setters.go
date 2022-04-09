package analyser

func (analyser *Analyser) setNewHalf() {
	analyser.halfs = append(analyser.halfs, &Half{
		ctName:      analyser.parser.GameState().TeamCounterTerrorists().ClanName(),
		tName:       analyser.parser.GameState().TeamTerrorists().ClanName(),
		halfCtScore: analyser.halfCtScore,
		halfTScore:  analyser.halfTScore,
	})
}

func (analyser *Analyser) setRoundEnd(tick int) {
	analyser.roundStarted = false
	analyser.currentRound.endTick = tick
	analyser.setRoundFinish()
}

func (analyser *Analyser) setRoundEndOfficial(tick int) {
	analyser.roundStarted = false
	analyser.currentRound.endOfficialTick = tick
	analyser.setRoundFinish()
}

func (analyser *Analyser) setMatchEnded() {
	analyser.matchEnded = true
}

func (analyser *Analyser) setRoundFinish() {
	analyser.rounds = append(analyser.rounds, analyser.currentRound)
	analyser.roundsPlayed++
	analyser.previousRound = analyser.currentRound
	analyser.currentRound = nil
}
