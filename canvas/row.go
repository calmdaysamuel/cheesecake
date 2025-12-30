package canvas

import "strings"

type Row []Cell

func (r Row) View() string {
	cells := make([]string, 0)
	for _, cell := range r {
		cells = append(cells, cell.View())
	}
	return strings.Join(cells, "")
}
