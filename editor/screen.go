package editor

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen
const lineNumberWidth = 4

func InitScreen() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	screen.Clear()
}

func RenderBuffer() {
	screen.Clear()

	lines := buffer.GetLines()

	for y, line := range lines {
		// Format line number
		lineNum := fmt.Sprintf("%*d ", lineNumberWidth-1, y+1)
		for x, r := range lineNum {
			screen.SetContent(x, y, r, nil, tcell.StyleDefault.Foreground(tcell.ColorGray))
		}

		// Draw the text line after the line number
		for x, r := range line {
			screen.SetContent(x+lineNumberWidth, y, r, nil, tcell.StyleDefault)
		}
	}

	// Draw mode or command input
	_, height := screen.Size()
	if inCommandMode {
		cmd := ":" + commandInput
		for i, r := range cmd {
			screen.SetContent(i, height-1, r, nil, tcell.StyleDefault)
		}
	} else {
		status := "-- " + ModeName() + " --"
		for i, r := range status {
			screen.SetContent(i, height-2, r, nil, tcell.StyleDefault)
		}
	}

	// Offset the cursor for line number
	screen.ShowCursor(buffer.CursorX+lineNumberWidth, buffer.CursorY)

	screen.Show()
	if statusMessage != "" {
	for i, r := range statusMessage {
		screen.SetContent(i, height-3, r, nil, tcell.StyleDefault)
	}
}
}



func CloseScreen() {
	if screen != nil {
		screen.Fini()
		fmt.Print("\033[2J\033[H") // ANSI: clear screen and move cursor to top-left
	}
}

func PollEvent() tcell.Event {
	return screen.PollEvent()
}

func SyncScreen() {
	screen.Sync()
}
