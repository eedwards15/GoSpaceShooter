package systems

var (
	WINDOWMANAGER *WindowManager
)

type WindowManager struct {
	SCREENWIDTH  int
	SCREENHEIGHT int
}

func NewWindowManager(width int, height int) *WindowManager {
	windowManger := WindowManager{
		SCREENWIDTH:  width,
		SCREENHEIGHT: height,
	}
	return &windowManger
}

func (windowManagerClass WindowManager) Center() (float64, float64) {
	return float64(windowManagerClass.SCREENWIDTH / 2), float64(windowManagerClass.SCREENHEIGHT / 2)
}
