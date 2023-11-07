package analyser

import (
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
)

func (analyser *Analyser) registerNetMessageHandlers() {
	if analyser.isSource2 {
		analyser.parser.RegisterNetMessageHandler(func(msg *msgs2.CSVCMsg_ServerInfo) {
			analyser.mapName = msg.GetMapName()
		})
	}
}
