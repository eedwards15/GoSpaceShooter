package definitions

type LevelDefinition struct {
	BackgroundMusic BackgroundMusic
}

type BackgroundMusic struct {
	Path       string
	SampleRate int
}
