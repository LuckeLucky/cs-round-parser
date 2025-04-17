package analyser

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newAnalyserByPath(path string) *Analyser {
	f, err := os.Open(path)

	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	an := NewAnalyser(f, false)
	an.FirstParse()
	f.Close()

	return an
}

func TestSource2Demo(t *testing.T) {
	an := newAnalyserByPath("../test-demos/main_cs2.dem")

	assert.Equal(t, "de_anubis", an.mapName)

	for _, players := range an.players {
		assert.Equal(t, 5, len(players))
	}
	assert.Equal(t, 0, len(an.spectators))

	assert.Equal(t, 3, an.rounds[11].ctScore)
	assert.Equal(t, 9, an.rounds[11].tScore)

	assert.Equal(t, 13, an.rounds[len(an.rounds)-1].ctScore)
	assert.Equal(t, 3, an.rounds[len(an.rounds)-1].tScore)

	assert.Equal(t, 16, an.roundsPlayed)
}
