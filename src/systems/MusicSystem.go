package systems

import (
	"SpaceShooter/src/definitions"
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

func NewMusicSystem(music definitions.BackgroundMusic) *MusicSystem {
	musicSystem := MusicSystem{}
	musicSystem.sampleRate = music.SampleRate
	f, _ := ebitenutil.OpenFile(music.Path)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(music.SampleRate, f)
	return &musicSystem
}

func (musicSystem *MusicSystem) LoadSong(music definitions.BackgroundMusic) *MusicSystem {
	musicSystem.player.Close()
	musicSystem.sampleRate = music.SampleRate
	f, _ := ebitenutil.OpenFile(music.Path)
	musicSystem.currentSong, _ = mp3.DecodeWithSampleRate(music.SampleRate, f)
	return musicSystem
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

func (musicSystem *MusicSystem) SetVolume(vol float64) {
	musicSystem.player.SetVolume(vol)
}
