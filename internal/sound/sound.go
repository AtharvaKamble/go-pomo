package sound

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"os"
)

const (
	SAMPLE_RATE = 44100
)

var (
	player   *AudioPlayer
	audioCtx *audio.Context
)

type AudioPlayer struct {
	SoundPath string
	Player    *audio.Player
	Ctx       *audio.Context
}

func init() {
	audioCtx = audio.NewContext(SAMPLE_RATE)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// GetPlayerInstance creates only a single player instance
func GetPlayerInstance(path string) *AudioPlayer {

	f, err := os.Open(path)
	check(err)

	stream, err := wav.DecodeWithoutResampling(f)
	check(err)

	p, err := audioCtx.NewPlayer(stream)
	check(err)

	return &AudioPlayer{SoundPath: path, Player: p, Ctx: audioCtx}
}

func (a *AudioPlayer) PlaySound() {
	err := a.Player.Rewind()
	check(err)

	a.Player.Play()

	//err = audio.Player.Close()
	//check(err)
}

// player = &Player{SoundPath: "/absolute/or/relative/path"}
