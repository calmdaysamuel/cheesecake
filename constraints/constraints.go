package constraints

import (
	"fmt"
)

type Constraints struct {
	MaxHeight int
	MaxWidth  int
}

func (c Constraints) String() string {
	return fmt.Sprintf("MaxHeight: %d, MaxWidth: %d", c.MaxHeight, c.MaxWidth)
}
func Tight(width, height int) Constraints {
	return Constraints{
		MaxHeight: height,
		MaxWidth:  width,
	}
}
