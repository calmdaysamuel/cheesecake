package canvas

// Partition partitions a slice into sub-slices of a specified size.
func Partition(slice []Cell, chunkSize int) Canvas {
	if chunkSize == 0 {
		return Canvas{}
	}
	if len(slice) == 0 {
		return Canvas{}
	}
	var chunks Canvas
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
