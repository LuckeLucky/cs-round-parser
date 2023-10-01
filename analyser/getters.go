package analyser

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

func (analyser *Analyser) getGameTick() (int, bool) {
	var err bool
	tick := analyser.parser.GameState().IngameTick()
	if tick < 0 {
		err = true
	}
	return tick, err
}

func (analyser *Analyser) getTeamName(ts *common.TeamState) string {
	name := ts.ClanName()
	if name == "" {
		sum := 0
		for _, p := range ts.Members() {
			sum += int(p.SteamID32())
		}
		return strconv.Itoa(sum)
	}
	return name
}
