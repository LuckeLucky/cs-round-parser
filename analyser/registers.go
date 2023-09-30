package analyser

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msg"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
)

func (analyser *Analyser) registerNetMessageHandlers() {
	// Register handler for net messages updates
	analyser.parser.RegisterNetMessageHandler(func(msg *msg.CNETMsg_SetConVar) {
		for _, cvar := range msg.Convars.Cvars {
			cvarName := cvar.GetName()
			cvarValue := cvar.GetValue()
			if cvarName == "mp_overtime_maxrounds" {
				analyser.overtimeMaxRounds, _ = strconv.Atoi(cvarValue)
			} else if cvarName == "mp_startmoney" {
				analyser.currentStartMoney, _ = strconv.Atoi(cvarValue)
				analyser.isMoneySet = true
			} else if cvarName == "mp_free_armor" {
				analyser.freeArmor, _ = strconv.Atoi(cvarValue)
			} else if cvarName == "mp_overtime_startmoney" {
				/*sometimes mp_overtime_startmoney is used instead of start_money for overtimes*/
				analyser.currentOvertimeStartMoney, _ = strconv.Atoi(cvarValue)
				analyser.isOvertimeMoneySet = true
			} else if cvarName == "mp_maxrounds" {
				analyser.maxRounds, _ = strconv.Atoi(cvarValue)
			}
		}
	})

	analyser.parser.RegisterNetMessageHandler(func(msg *msgs2.CSVCMsg_ServerInfo) {
		analyser.mapName = msg.GetMapName()
	})

	analyser.parser.RegisterNetMessageHandler(func(msg *msgs2.CNETMsg_SetConVar) {
		for _, cvar := range msg.Convars.Cvars {
			cvarName := cvar.GetName()
			cvarValue := cvar.GetValue()
			if cvarName == "mp_maxrounds" {
				analyser.maxRounds, _ = strconv.Atoi(cvarValue)
			}
		}
	})
}

func (analyser *Analyser) registerMatchEventHandlers() {
	//Round start
	analyser.parser.RegisterEventHandler(func(e events.RoundStart) { analyser.handlerRoundStart(e) })
	analyser.parser.RegisterEventHandler(func(e events.MatchStartedChanged) { analyser.handlerRoundStart(e) })
	analyser.parser.RegisterEventHandler(func(e events.RoundFreezetimeEnd) { analyser.handlerRoundStart(e) })
	//Round ends
	analyser.parser.RegisterEventHandler(func(e events.RoundEnd) { analyser.handlerRoundEnd(e) })
	analyser.parser.RegisterEventHandler(func(e events.RoundEndOfficial) { analyser.handlerRoundEndOfficial(e) })

	//Handle side switch
	analyser.parser.RegisterEventHandler(func(e events.TeamSideSwitch) { analyser.handlerSideSwitch() })
}
