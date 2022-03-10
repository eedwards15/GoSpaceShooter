package systems

var (
	WINDOWMANAGER *WindowManager
)

type WindowManager struct {
	SCREENWIDTH  int
	SCREENHEIGHT int
}

func InitWindowManager(width int, height int) {
	if WINDOWMANAGER == nil {
		windowManger := &WindowManager{
			SCREENWIDTH:  width,
			SCREENHEIGHT: height,
		}
		WINDOWMANAGER = windowManger
	}
}

func (windowManagerClass WindowManager) Center() (float64, float64) {
	return float64(windowManagerClass.SCREENWIDTH / 2), float64(windowManagerClass.SCREENHEIGHT / 2)
}
