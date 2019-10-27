package bmh

func Search(buffer, pattern []byte) (position int) {
	var table = Table{}
	var idxBuffer int
	var idxPattern int
	var lenPattern = len(pattern)
	var lenBuffer = len(buffer)

	position = -1
	table.Build(pattern)
	for (idxBuffer + lenPattern) <= lenBuffer {
		idxPattern = lenPattern  -1

		for buffer[idxBuffer + idxPattern] == pattern[idxPattern] {
			if idxPattern == 0 {
				position = idxBuffer
				return position
			}
			idxPattern -= 1
		}

		idxBuffer += table.Content[buffer[idxBuffer + (lenPattern - 1)]]
	}

	return position
}