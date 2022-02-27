package scenes

import "SpaceShooter/src/interfaces"

type SceneManager []interfaces.IScene

func (s SceneManager) Peek() *interfaces.IScene {
	l := len(s)
	return &s[l-1]
}

func (s SceneManager) Push(v interfaces.IScene) []interfaces.IScene {
	return append(s, v)
}

func (s SceneManager) Pop() ([]interfaces.IScene, *interfaces.IScene) {
	if len(s) <= 0 {
		return make([]interfaces.IScene, 0), nil
	}

	l := len(s)

	return s[:l-1], &s[l-1]
}
