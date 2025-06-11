package editor

type EditorMode int

const (
	NormalMode EditorMode = iota
	InsertMode
)

var currentMode = NormalMode

func GetMode() EditorMode {
	return currentMode
}

func SetMode(mode EditorMode) {
	currentMode = mode
}

func ModeName() string {
	switch currentMode {
	case NormalMode:
		return "NORMAL"
	case InsertMode:
		return "INSERT"
	default:
		return "UNKNOWN"
	}
}
