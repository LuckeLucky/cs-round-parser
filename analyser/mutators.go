package analyser

func (analyser *Analyser) setDefault() {
	analyser.spectators = make(map[uint64]string)
	analyser.players = make(map[string]map[uint64]string)
}
