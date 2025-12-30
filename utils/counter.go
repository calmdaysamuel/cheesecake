package utils

func TrueCounter(bools ...bool) (total int) {
	for _, b := range bools {
		if b {
			total++
		}
	}
	return total
}
