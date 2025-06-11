package editor

import ( "os"
		"strings"
		"bufio"
		"fmt"
		"github.com/gdamore/tcell/v2")

var (commandInput string
 	inCommandMode bool
 	statusMessage string 
 	currentFilename string)

func StartEditor(filename string) {
	InitScreen()
	defer CloseScreen()

	if filename != "" {
		loadFile(filename)
	}
	currentFilename = filename

	RenderBuffer()

	for {
		ev := PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch GetMode() {

			case NormalMode:
				if inCommandMode {
					switch ev.Key() {
					case tcell.KeyEnter:
						handleCommand(commandInput)
						commandInput = ""
						inCommandMode = false
						RenderBuffer()
					case tcell.KeyBackspace, tcell.KeyBackspace2:
						if len(commandInput) > 0 {
							commandInput = commandInput[:len(commandInput)-1]
						}
						RenderBuffer()
					default:
						if ev.Rune() != 0 {
							commandInput += string(ev.Rune())
						}
						RenderBuffer()
					}
				} else {
					switch ev.Rune() {
					case 'i':
						SetMode(InsertMode)
					case 'h':
						buffer.MoveCursor(-1, 0)
					case 'l':
						buffer.MoveCursor(1, 0)
					case 'j':
						buffer.MoveCursor(0, 1)
					case 'k':
						buffer.MoveCursor(0, -1)
					case ':':
						inCommandMode = true
						commandInput = ""
					}
					RenderBuffer()
				}

			case InsertMode:
				if ev.Key() == tcell.KeyEsc {
					SetMode(NormalMode)
				} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
					buffer.Backspace()
				} else if ev.Key() == tcell.KeyEnter {
					buffer.NewLine()
				} else if ev.Rune() != 0 {
					buffer.InsertRune(ev.Rune())
				}
				RenderBuffer()
		}

		case *tcell.EventResize:
			SyncScreen()
			RenderBuffer()
		}
	}
}



func handleCommand(cmd string) {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "q":
		CloseScreen()
		os.Exit(0)

	case "w":
		if len(args) < 2 {
			return
		}
		filename := args[1]
		writeToFile(filename)

	case "wq":
		if len(args) >= 2 {
			currentFilename = args[1]
		}
		if currentFilename != "" {
			writeToFile(currentFilename)
			CloseScreen()
			os.Exit(0)
		} else {
			statusMessage = "No filename specified"
		}
}
	
}

func loadFile(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		statusMessage = fmt.Sprintf("\"%s\" [New File]", filename)
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}
	buffer.SetLines(lines)
}

func writeToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		statusMessage = "Error writing to file"
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	var bytes, lines int
	for _, line := range buffer.GetLines() {
		for _, r := range line {
			n, _ := writer.WriteRune(r)
			bytes += n
		}
		writer.WriteRune('\n')
		bytes++
		lines++
	}

	writer.Flush() // VERY IMPORTANT

	statusMessage = fmt.Sprintf("\"%s\" %dL, %dB written", filename, lines, bytes)
}