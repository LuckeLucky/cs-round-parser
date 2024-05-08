package analyser

import (
	"io"

	"github.com/LuckeLucky/cs-round-parser/utils"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
)

type Analyser struct {
	parser    demoinfocs.Parser
	isSource2 bool

	cfg     demoinfocs.ParserConfig
	mapName string

	rounds        []*Round
	currentRound  *Round
	previousRound *Round
	roundsPlayed  int

	roundStarted bool
	matchEnded   bool

	//Current ScoreBoard scores
	ctScore int
	tScore  int

	//Convars -----------------
	maxRounds         int
	overtimeMaxRounds int
	freeArmor         int

	//Demo participants
	spectators map[uint64]string
	players    map[string](map[uint64]string)

	isSimpleReader bool
}

func NewAnalyser(demostream io.Reader, isSimpleReader bool) *Analyser {
	analyser := &Analyser{}
	analyser.cfg = demoinfocs.DefaultParserConfig

	parser := demoinfocs.NewParserWithConfig(demostream, analyser.cfg)
	analyser.parser = parser
	analyser.isSimpleReader = isSimpleReader

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
	if !analyser.isSimpleReader {
		analyser.printRoundsPlayed()
	}
	analyser.printPlayers()
	analyser.printSpectators()
}
