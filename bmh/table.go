package bmh

type Table struct {
	Content []int
}

// We initialize the table with the length of the pattern
// When we find an unknown character in the pattern, the number will be the length of the pattern
// num = the length of the pattern
func (t *Table) initContent(num int) {
	t.Content = make([]int, 256)

	for i := 0; i < 256; i++ {
		t.Content[i] = num
	}
}

func (t *Table) Build(pattern []byte) {
	var lenPattern = len(pattern)

	t.initContent(lenPattern)
	for i := 0; i < (lenPattern - 1); i++ { // lenPattern - 1 because the last is lenPattern that set in "initContent"
		// When we find a character that exists in the table, we return the number of jumps
		// The number of jumps is the position of the character from the end
		// If a character is several times, we take the last one
		t.Content[pattern[i]] = (lenPattern - 1) - i
	}
}