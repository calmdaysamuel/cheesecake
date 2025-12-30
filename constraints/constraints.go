package constraints

type Constraints struct {
	MaxHeight int
	MaxWidth  int
}

func Tight(width, height int) Constraints {
	return Constraints{
		MaxHeight: height,
		MaxWidth:  width,
	}
}
