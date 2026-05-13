//Let's write some crazy code ...

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/RudraPratapDev/NoteX/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.InitializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error %v", err)
		os.Exit(1)
	}
}
