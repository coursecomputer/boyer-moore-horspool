package bmh

type Table struct {
	Content [256]int
}

func (t *Table) initContent(num int) {
	for i := 0; i < 256; i++ {
		t.Content[i] = num
	}
}

func (t *Table) Build(pattern []byte) {
	var lenPattern = len(pattern)

	t.initContent(lenPattern)
	for i := 0; i < (lenPattern - 1); i++ {
		t.Content[pattern[i]] = (lenPattern - 1) - i
	}
}