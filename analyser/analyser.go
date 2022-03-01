package analyser

import (
	"fmt"
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

	rounds       []*Round
	currentRound *Round
	roundsPlayed int
	halfs        []*Half
	roundStarted bool

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

func (analyser *Analyser) SimpleRun() {
	analyser.handleHeader()
	analyser.setDefault()

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()

	var err error
	for ok := true; ok; ok, err = analyser.parser.ParseNextFrame() {
		utils.CheckError(err)
	}

	analyser.printHalfs()
	analyser.printMap()
	fmt.Printf("Rounds played:%d\n", analyser.roundsPlayed)
}
