package main

import ( 
	"os"
	// "fmt"	
	"github.com/adhit-420/govim/editor"
)


func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]

		// Try to create file if it doesn't exist
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// Create the file (empty, so user can :w later)
			empty, err := os.Create(filename)
			if err != nil {
				panic("Failed to create file: " + err.Error())
			}
			empty.Close()
		}
	}

	editor.StartEditor(filename)
}
