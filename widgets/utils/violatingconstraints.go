package utils

import (
	"context"
	"github.com/calmdaysamuel/cheesecake/canvas"
	"github.com/calmdaysamuel/cheesecake/constraints"
	"github.com/calmdaysamuel/cheesecake/size"
	"github.com/calmdaysamuel/cheesecake/widgets/colors"
)

func ViolatingConstraintsCanvas(ctx context.Context, box constraints.Constraints, c canvas.Canvas) (canvas.Canvas, size.Size, error) {
	s := canvas.Size(c)
	isOverflowing, direction := box.ViolatesConstraints(s)
	if !isOverflowing {
		return c, s, nil
	}
	c = canvas.Truncate(c, box.MaxWidth, box.MaxHeight)
	if direction == constraints.Both || direction == constraints.Vertically {
		c = OverflowingCanvasVertical(c)
	}
	if direction == constraints.Both || direction == constraints.Horizontally {
		c = OverflowingCanvasHorizontal(c)
	}
	return c, canvas.Size(c), nil
}

func OverflowingCanvasVertical(c canvas.Canvas) canvas.Canvas {
	i := len(c) - 1
	for j, cell := range c[i] {
		bgColor := colors.Yellow
		if j%2 != i%2 {
			bgColor = colors.Black
		}
		temp := cell
		temp.BgColor = bgColor
		temp.Bold = true
		temp.Transparent = false
		temp.Runes = []rune("x")
		c[i][j] = temp
	}
	return c
}

func OverflowingCanvasHorizontal(c canvas.Canvas) canvas.Canvas {
	for i := 0; i < len(c); i++ {
		if len(c[i]) <= 0 {
			continue
		}
		j := len(c[i]) - 1
		bgColor := colors.Yellow
		if j%2 != i%2 {
			bgColor = colors.Black
		}
		temp := c[i][j]
		temp.BgColor = bgColor
		temp.Bold = true
		temp.Transparent = false
		temp.Runes = []rune("x")
		c[i][j] = temp
	}
	return c
}
