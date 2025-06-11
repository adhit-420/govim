package editor

type Buffer struct {
	lines   [][]rune
	CursorX int
	CursorY int
}



var buffer = NewBuffer()

func NewBuffer() *Buffer {
	return &Buffer{
		lines:   [][]rune{{}}, // Start with one empty line
		CursorX: 0,
		CursorY: 0,
	}
}

func (b *Buffer) Backspace() {
	if b.CursorX == 0 {
		if b.CursorY == 0 {
			return
		}
		// Merge current line into previous
		prevLineLen := len(b.lines[b.CursorY-1])
		b.lines[b.CursorY-1] = append(b.lines[b.CursorY-1], b.lines[b.CursorY]...)
		b.lines = append(b.lines[:b.CursorY], b.lines[b.CursorY+1:]...)
		b.CursorY--
		b.CursorX = prevLineLen
	} else {
		// Remove character before cursor
		line := b.lines[b.CursorY]
		b.lines[b.CursorY] = append(line[:b.CursorX-1], line[b.CursorX:]...)
		b.CursorX--
	}
}

// Handle Enter (newline)
func (b *Buffer) NewLine() {
	line := b.lines[b.CursorY]
	newLine := line[b.CursorX:]
	b.lines[b.CursorY] = line[:b.CursorX]
	b.lines = append(b.lines[:b.CursorY+1], append([][]rune{{}}, b.lines[b.CursorY+1:]...)...)
	b.lines[b.CursorY+1] = newLine
	b.CursorY++
	b.CursorX = 0
}

// Insert character at current position
func (b *Buffer) InsertRune(r rune) {

	if len(b.lines) == 0 {
		b.lines = append(b.lines, []rune{})
	}

	if b.CursorY >= len(b.lines) {
		for len(b.lines) <= b.CursorY {
			b.lines = append(b.lines, []rune{})
		}
	}

	line := b.lines[b.CursorY]
	CursorX := b.CursorX

	if CursorX >= len(line) {
		line = append(line, r)
	} 
	
	line = append(line[:CursorX], append([]rune{r}, line[CursorX:]...)...)
	b.lines[b.CursorY] = line
	b.CursorX++
}

// Move cursor within bounds
func (b *Buffer) MoveCursor(dx, dy int) {
	b.CursorY += dy
	if b.CursorY < 0 {
		b.CursorY = 0
	} else if b.CursorY >= len(b.lines) {
		b.CursorY = len(b.lines) - 1
	}

	b.CursorX += dx
	if b.CursorX < 0 {
		b.CursorX = 0
	}
	lineLen := len(b.lines[b.CursorY])
	if b.CursorX > lineLen {
		b.CursorX = lineLen
	}
}

// Return lines for drawing
func (b *Buffer) GetLines() [][]rune {
	return b.lines
}

func (b *Buffer) SetLines(lines [][]rune) {
	b.lines = lines
	b.CursorX = 0
	b.CursorY = 0
}