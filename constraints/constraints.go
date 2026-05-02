package constraints

import (
	"errors"
	"fmt"
	"github.com/calmdaysamuel/cheesecake/size"
	"math"
)

type ViolatingConstraints = int

const (
	Vertically   = 1
	Horizontally = 2
	Both         = 3
)

var Max = Constraints{
	MaxHeight: math.MaxInt,
	MaxWidth:  math.MaxInt,
}

type Constraints struct {
	MaxHeight int `json:"maxHeight"`
	MaxWidth  int `json:"maxWidth"`
}

func (c Constraints) String() string {
	return fmt.Sprintf("MaxHeight: %d, MaxWidth: %d", c.MaxHeight, c.MaxWidth)
}

func (c Constraints) Validate() error {
	if c.MaxWidth < 0 {
		return errors.New("invalid constraints: width is less than zero")
	}

	if c.MaxHeight < 0 {
		return errors.New("invalid constraints: height is less than zero")
	}
	return nil
}

func (c Constraints) IsZero() bool {
	if c.MaxWidth == 0 || c.MaxHeight == 0 {
		return true
	}

	return false
}

func (c Constraints) ViolatesConstraints(size size.Size) (isOverflowing bool, direction ViolatingConstraints) {
	if size.Height > c.MaxHeight && size.Width > c.MaxWidth {
		return true, Both
	}
	if size.Height > c.MaxHeight {
		return true, Vertically
	}
	if size.Width > c.MaxWidth {
		return true, Horizontally
	}
	return false, 0
}
func Tight(width, height int) Constraints {
	return Constraints{
		MaxHeight: height,
		MaxWidth:  width,
	}
}
