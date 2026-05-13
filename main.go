//Let's write some crazy code ...

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	newFIleInput     textinput.Model
	fileInputVisible bool
}

func (m model) Init() tea.Cmd {

	return nil

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	//Is it a key pressed?
	case tea.KeyMsg:
		//which key was pressed?
		switch msg.String() {

		// Oh , so you want to Quit.
		case "ctrl+c", "q":
			return m, tea.Quit

		//Create a new file.
		case "ctrl+n":
			m.fileInputVisible = true
			return m, nil

		case "enter":
			//todo create file
			return m, nil

		}
	}
	if m.fileInputVisible {
		m.newFIleInput, cmd = m.newFIleInput.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {

	var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("16")).
		PaddingTop(2).
		PaddingLeft(4).
		PaddingRight(4).PaddingBottom(2)

	welcome := titleStyle.Render("Welcome to NoteX 📝")

	var helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Background(lipgloss.Color("16")).
		PaddingLeft(4).
		PaddingRight(4).
		MarginTop(1)

	keyHelp := helpStyle.Render(
		"Ctrl+N: new file • Ctrl+L: list • Esc: back/save • Ctrl+S: save • Ctrl+Q: quit • Ctrl+P: export to PDF",
	)

	view := ""
	if m.fileInputVisible {

		inputBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			MarginTop(1).
			Width(40)

		view = inputBox.Render(m.newFIleInput.View())

	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, keyHelp)

}

func initializeModel() model {

	// inititalize new file input
	ti := textinput.New()
	ti.Placeholder = "Enter file name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	// Styling
	ti.Prompt = "📄 "

	ti.TextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)

	ti.PromptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)

	ti.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)

	ti.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("212"))

	return model{
		newFIleInput:     ti,
		fileInputVisible: false,
	}

}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {

		fmt.Printf("error %v", err)
		os.Exit(1)
	}
}
