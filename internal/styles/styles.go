package styles

import "github.com/charmbracelet/lipgloss"

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("205")).
	Background(lipgloss.Color("235")).
	PaddingLeft(2).
	PaddingRight(2).
	PaddingTop(0).
	PaddingBottom(0)

var HeaderDimStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241")).
	Background(lipgloss.Color("235")).
	PaddingRight(2)

var StatusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("243")).
	Background(lipgloss.Color("235")).
	PaddingLeft(2).
	PaddingRight(2)

var StatusBarAccent = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Background(lipgloss.Color("235")).
	Bold(true)

var InputBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(1, 2).
	MarginTop(2).
	Width(44)

var InputLabelStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("86")).
	Bold(true).
	MarginTop(2).
	PaddingLeft(1)

var HomeHintStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("238")).
	Italic(true).
	PaddingLeft(2).
	MarginTop(1)

var TextInputTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true)

var TextInputPromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("86")).
	Bold(true)

var TextInputPlaceholderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	Italic(true)

var TextInputCursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("212"))
