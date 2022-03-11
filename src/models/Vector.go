package models

type Vector2 struct {
	Xpos float64
	Ypos float64
}

func NewVector2(xpos, ypos float64) *Vector2 {
	v := &Vector2{
		Xpos: xpos,
		Ypos: ypos,
	}
	return v
}
func (vectorClass *Vector2) SetX(xpos float64) {
	vectorClass.Xpos = xpos
}

func (vectorClass *Vector2) SetY(ypos float64) {
	vectorClass.Ypos = ypos
}
