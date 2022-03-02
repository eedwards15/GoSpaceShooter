package systems

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	MUSICSYSTEM *MusicSystem
)

type MusicSystem struct {
	currentSong *mp3.Stream
	sampleRate  int
	player      *audio.Player
}

func NewMusicSystem(songLocation string, sampleRate int) *MusicSystem {
	musicSystem := MusicSystem{}
	musicSystem.sampleRate = sampleRate
	f, _ := ebitenutil.OpenFile(songLocation)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(sampleRate, f)
	return &musicSystem
}

func (musicSystem *MusicSystem) LoadSong(songLocation string, sampleRate int) {
	musicSystem.player.Close()
	musicSystem.sampleRate = sampleRate
	f, _ := ebitenutil.OpenFile(songLocation)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(sampleRate, f)
}

func (musicSystem *MusicSystem) PlaySong() {
	if audio.CurrentContext() == nil {
		audioContext := audio.NewContext(musicSystem.sampleRate)
		loop := audio.NewInfiniteLoop(musicSystem.currentSong, musicSystem.currentSong.Length())
		musicSystem.player, _ = audioContext.NewPlayer(loop)
		musicSystem.player.Play()
		return
	}

	loop := audio.NewInfiniteLoop(musicSystem.currentSong, musicSystem.currentSong.Length())
	musicSystem.player, _ = audio.CurrentContext().NewPlayer(loop)
	musicSystem.player.Play()
}

func (musicSystem *MusicSystem) Pause() {
	musicSystem.player.Pause()
}

func (musicSystem *MusicSystem) Rewind() {
	musicSystem.player.Rewind()
}
