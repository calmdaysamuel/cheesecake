package constraints

import (
	"errors"
	"fmt"
)

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
func Tight(width, height int) Constraints {
	return Constraints{
		MaxHeight: height,
		MaxWidth:  width,
	}
}
