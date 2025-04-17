package analyser

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/msg"
)

func (analyser *Analyser) registerNetMessageHandlers() {
	analyser.parser.RegisterNetMessageHandler(func(msg *msg.CSVCMsg_ServerInfo) {
		analyser.mapName = msg.GetMapName()
	})
	analyser.parser.RegisterNetMessageHandler(func(msg *msg.CNETMsg_SetConVar) {
		for _, cvar := range msg.Convars.Cvars {
			analyser.handleSetConvar(cvar.GetName(), cvar.GetValue())
		}
	})
}

func (analyser *Analyser) handleSetConvar(name string, value string) {
	switch name {
	case "mp_overtime_maxrounds":
		analyser.overtimeMaxRounds, _ = strconv.Atoi(value)
	case "mp_free_armor":
		analyser.freeArmor, _ = strconv.Atoi(value)
	case "mp_maxrounds":
		analyser.maxRounds, _ = strconv.Atoi(value)
	}
}
