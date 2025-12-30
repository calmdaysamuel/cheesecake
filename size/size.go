package size

type Size struct {
	Height int
	Width  int
}

func (e *Size) GetSize() Size {
	return *e
}
