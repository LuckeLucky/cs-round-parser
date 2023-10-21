package analyser

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msg"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
)

func (analyser *Analyser) registerNetMessageHandlers() {
	// Register handler for net messages updates
	analyser.parser.RegisterNetMessageHandler(func(msg *msg.CNETMsg_SetConVar) {
		for _, cvar := range msg.Convars.Cvars {
			analyser.handleSetConvar(cvar.GetName(), cvar.GetValue())
		}
	})

	if analyser.isSource2 {
		analyser.parser.RegisterNetMessageHandler(func(msg *msgs2.CSVCMsg_ServerInfo) {
			analyser.mapName = msg.GetMapName()
		})
		analyser.parser.RegisterNetMessageHandler(func(msg *msgs2.CNETMsg_SetConVar) {
			for _, cvar := range msg.Convars.Cvars {
				analyser.handleSetConvar(cvar.GetName(), cvar.GetValue())
			}
		})
	}
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
