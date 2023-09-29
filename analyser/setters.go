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

func (analyser *Analyser) setParticipants() {
	for _, participant := range analyser.parser.GameState().Participants().Connected() {
		//1 corresponds to team Spectators
		steamID := participant.SteamID64
		name := participant.Name
		if participant.Team == 1 {
			if _, ok := analyser.spectators[steamID]; !ok {
				analyser.spectators[steamID] = name
			}
		} else if participant.Team != 0 {
			clanName := participant.TeamState.ClanName()
			if _, ok := analyser.players[clanName]; !ok {
				analyser.players[clanName] = make(map[uint64]string)
			}
			if _, ok := analyser.players[clanName][steamID]; !ok {
				analyser.players[clanName][steamID] = name
			}
		}
	}
}