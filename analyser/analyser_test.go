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

	an := NewAnalyser(f)
	an.FirstParse()
	f.Close()

	return an
}

func TestSourceDemo(t *testing.T) {
	an := newAnalyserByPath("../test-demos/main_cs.dem")

	assert.Equal(t, false, an.isSource2)
	assert.Equal(t, "de_mirage", an.mapName)

	assert.Equal(t, 5, an.rounds[14].ctScore)
	assert.Equal(t, 10, an.rounds[14].tScore)

	assert.Equal(t, 16, an.rounds[len(an.rounds)-1].ctScore)
	assert.Equal(t, 6, an.rounds[len(an.rounds)-1].tScore)

	for _, players := range an.players {
		assert.Equal(t, 5, len(players))
	}
	assert.Equal(t, 2, len(an.spectators))

	assert.Equal(t, 22, an.roundsPlayed)
}

func TestSource2Demo(t *testing.T) {
	an := newAnalyserByPath("../test-demos/main_cs2.dem")

	assert.Equal(t, true, an.isSource2)
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

func TestOvertimes(t *testing.T) {
	an := newAnalyserByPath("../test-demos/cs_overtime.dem")

	assert.Equal(t, 15, an.rounds[32].ctScore)
	assert.Equal(t, 18, an.rounds[32].tScore)
}

func TestNoDataPacket(t *testing.T) {
	an := newAnalyserByPath("../test-demos/esportal_no_data.dem")
	assert.Equal(t, 23, an.roundsPlayed)
}
