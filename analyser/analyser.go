package analyser

import (
	"io"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
)

type Analyser struct {
	parser demoinfocs.Parser

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

// Used to gather information about RoundStart..End and team scores
func (analyser *Analyser) FirstParse() {
	analyser.setDefault()

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()

	// Parse to end
	err := analyser.parser.ParseToEnd()
	if err != demoinfocs.ErrUnexpectedEndOfDemo {
		analyser.printEndOfParsing()
	}

	analyser.printMap()
	if !analyser.isSimpleReader {
		analyser.printRoundsPlayed()
	}
	analyser.printPlayers()
	analyser.printSpectators()
}
