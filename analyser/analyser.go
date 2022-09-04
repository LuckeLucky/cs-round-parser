package analyser

import (
	"io"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/gogo/protobuf/proto"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

type Analyser struct {
	parser demoinfocs.Parser

	cfg     demoinfocs.ParserConfig
	mapName string

	rounds        []*Round
	currentRound  *Round
	previousRound *Round
	roundsPlayed  int
	halfs         []*Half

	roundStarted bool

	matchEnded bool

	//Current ScoreBoard scores
	ctScore int
	tScore  int
	//half scores
	halfCtScore int
	halfTScore  int

	//Convars -----------------
	isMoneySet                bool
	isOvertimeMoneySet        bool
	currentStartMoney         int
	currentOvertimeStartMoney int
	overtimeMaxRounds         int
	freeArmor                 int

	//participants
	spectators map[uint64]string
}

func NewAnalyser(demostream io.Reader) *Analyser {
	analyser := &Analyser{}
	analyser.cfg = demoinfocs.DefaultParserConfig
	analyser.cfg.AdditionalNetMessageCreators = map[int]demoinfocs.NetMessageCreator{
		6: func() proto.Message {
			return new(msg.CNETMsg_SetConVar)
		},
	}

	parser := demoinfocs.NewParserWithConfig(demostream, analyser.cfg)
	analyser.parser = parser

	return analyser

}

func (analyser *Analyser) handleHeader() {
	header, err := analyser.parser.ParseHeader()
	utils.CheckError(err)
	analyser.mapName = header.MapName
}

// Used to gather information about RoundStart..End and team scores
func (analyser *Analyser) FirstParse() {
	analyser.handleHeader()
	analyser.setDefault()

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()
	analyser.registerPlayerEventHandlers()

	// Parse to end
	err := analyser.parser.ParseToEnd()
	if err != demoinfocs.ErrUnexpectedEndOfDemo {
		utils.CheckError(err)
	}

	analyser.printHalfs()
	analyser.printMap()
	analyser.printRoundsPlayed()
	analyser.printSpectators()
}
