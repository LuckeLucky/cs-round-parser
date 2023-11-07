package analyser

import (
	"io"

	"github.com/LuckeLucky/cs-round-parser/utils"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
)

type Analyser struct {
	parser    demoinfocs.Parser
	isSource2 bool
	mapName   string

	//Demo participants
	spectators map[uint64]string
	players    map[string](map[uint64]string)
}

func NewAnalyser(demostream io.Reader) *Analyser {
	analyser := &Analyser{}
	parser := demoinfocs.NewParserWithConfig(demostream, demoinfocs.DefaultParserConfig)
	analyser.parser = parser
	return analyser

}

func (analyser *Analyser) handleHeader() {
	header, err := analyser.parser.ParseHeader()
	utils.CheckError(err)
	analyser.mapName = header.MapName
	analyser.isSource2 = header.Filestamp == "PBDEMS2"
}

// Used to gather information about RoundStart..End and team scores
func (analyser *Analyser) FirstParse() {
	analyser.handleHeader()
	analyser.setDefault()

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()

	// Parse to end
	err := analyser.parser.ParseToEnd()
	if err != demoinfocs.ErrUnexpectedEndOfDemo {
		utils.CheckError(err)
	}

	analyser.printMap()
	analyser.printPlayers()
	analyser.printSpectators()
}
