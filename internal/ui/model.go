package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/RudraPratapDev/NoteX/internal/config"
	"github.com/RudraPratapDev/NoteX/internal/models"
	"github.com/RudraPratapDev/NoteX/internal/storage"
	"github.com/RudraPratapDev/NoteX/internal/styles"
)

type Model struct {
	newFIleInput     textinput.Model
	fileInputVisible bool
	fileListVisible  bool
	openedFromList   bool
	currentFile      *os.File
	noteTextArea     textarea.Model
	noteList         list.Model
	width            int
	height           int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := styles.DocStyle.GetFrameSize()
		m.noteList.SetSize(msg.Width-h, msg.Height-v-10)
		m.noteTextArea.SetWidth(msg.Width - 2)
		m.noteTextArea.SetHeight(msg.Height - 6)

	//Is it a key pressed?
	case tea.KeyMsg:
		//which key was pressed?
		switch msg.String() {

		// Oh , so you want to Quit.
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit

		//Create a new file.
		case "ctrl+n":
			m.fileInputVisible = true
			return m, nil

		case "ctrl+s":

			//write file from textArea to file
			if m.currentFile == nil {
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("file could not be saved :(")
				return m, nil
			}
			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("file could not be saved :(")
				return m, nil
			}
			if _, err := m.currentFile.WriteString(m.noteTextArea.Value()); err != nil {
				fmt.Println("file could not be saved :(")
				return m, nil
			}
			if err := m.currentFile.Close(); err != nil {
				fmt.Println("Some error occured .")
			}
			m.currentFile = nil
			m.noteTextArea.SetValue("")
			m.noteList.SetItems(storage.ListFiles())

			if m.openedFromList {
				m.openedFromList = false
				m.fileListVisible = true
			}

			return m, nil

		//Delete the selected or currently open note.
		case "ctrl+d":
			if m.currentFile != nil {
				fp := m.currentFile.Name()
				m.currentFile.Close()
				m.currentFile = nil
				m.noteTextArea.SetValue("")
				os.Remove(fp)
				m.noteList.SetItems(storage.ListFiles())
				if m.openedFromList {
					m.openedFromList = false
					m.fileListVisible = true
				}
				return m, nil
			}

			if m.fileListVisible {
				item, ok := m.noteList.SelectedItem().(models.Item)
				if ok {
					fp := fmt.Sprintf("%s/%s", config.VaultDir, item.Title_)
					os.Remove(fp)
					m.noteList.SetItems(storage.ListFiles())
				}
				return m, nil
			}

		case "ctrl+l":
			m.fileListVisible = true
			return m, nil

		case "esc":
			if m.fileInputVisible {
				m.fileInputVisible = false
				return m, nil
			}

			if m.currentFile != nil {
				m.noteTextArea.SetValue("")
				m.currentFile = nil
				if m.openedFromList {
					m.openedFromList = false
					m.fileListVisible = true
				}
				return m, nil
			}

			if m.fileListVisible {
				if m.noteList.FilterState() == list.Filtering {
					break
				}
				m.fileListVisible = false
			}

			return m, nil

		case "enter":

			//if showing list
			if m.fileListVisible {
				item, ok := m.noteList.SelectedItem().(models.Item)
				if ok {
					fp := fmt.Sprintf("%s/%s", config.VaultDir, item.Title_)
					content, err := os.ReadFile(fp)
					if err != nil {
						fmt.Println("Error fetching the selected note")
						return m, nil
					}
					m.noteTextArea.SetValue(string(content))
					f, err := os.OpenFile(fp, os.O_RDWR, 0644)
					if err != nil {
						fmt.Println("Error fetching the selected note")
						return m, nil
					}
					m.currentFile = f
					m.fileListVisible = false
					m.openedFromList = true
				}
				return m, nil
			}

			//only take enter case when input field visible
			if !m.fileInputVisible {
				break
			}
			//todo create file
			filename := m.newFIleInput.Value()
			if filename != "" {
				fp := fmt.Sprintf("%s/%s.md", config.VaultDir, filename)

				//to prevent overwritting of existing file if any
				_, err := os.Stat(fp)
				if err == nil {
					return m, nil
				}
				f, err := os.Create(fp)
				if err != nil {
					log.Fatal(err)
				}
				m.currentFile = f
				m.fileInputVisible = false
				m.newFIleInput.SetValue("")
			}
			return m, nil

		}
	}
	if m.fileInputVisible {
		m.newFIleInput, cmd = m.newFIleInput.Update(msg)
	}
	if m.currentFile != nil {
		m.noteTextArea, cmd = m.noteTextArea.Update(msg)
	}
	if m.fileListVisible {
		m.noteList, cmd = m.noteList.Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	title := styles.HeaderStyle.Render(" NoteX 📝 ")
	subtitle := styles.HeaderDimStyle.Render("terminal notepad")
	header := title + subtitle

	var content string
	switch {
	case m.fileInputVisible:
		label := styles.InputLabelStyle.Render("New note name:")
		box := styles.InputBoxStyle.Render(m.newFIleInput.View())
		content = label + "\n" + box

	case m.currentFile != nil:
		content = m.noteTextArea.View()

	case m.fileListVisible:
		content = m.noteList.View()

	default:
		content = styles.HomeHintStyle.Render("Ctrl+N to create a note  •  Ctrl+L to open a note")
	}

	var status string
	switch {
	case m.currentFile != nil:
		status = styles.StatusBarStyle.Render("Ctrl+S save  •  Ctrl+D delete  •  Esc discard  •  Ctrl+C quit")
	case m.fileInputVisible:
		status = styles.StatusBarStyle.Render("Enter confirm  •  Esc cancel  •  Ctrl+C quit")
	case m.fileListVisible:
		status = styles.StatusBarStyle.Render("↑↓ navigate  •  Enter open  •  Ctrl+D delete  •  Esc back  •  Ctrl+C quit")
	default:
		status = styles.StatusBarStyle.Render("Ctrl+N new  •  Ctrl+L list  •  Ctrl+C quit")
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", header, content, status)
}

func InitializeModel() Model {
	//init the folder for storing notes
	err := os.MkdirAll(config.VaultDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	//initialize new text area
	ta := textarea.New()
	ta.Placeholder = "Jot down anything…"
	ta.Focus()
	ta.ShowLineNumbers = false

	// inititalize new file input
	ti := textinput.New()
	ti.Placeholder = "Enter file name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	//pulling list
	ist := storage.ListFiles()
	noteList := list.New(ist, list.NewDefaultDelegate(), 0, 0)
	noteList.Title = "List of Notes"

	// Styling
	ti.Prompt = "📄 "
	ti.TextStyle = styles.TextInputTextStyle
	ti.PromptStyle = styles.TextInputPromptStyle
	ti.PlaceholderStyle = styles.TextInputPlaceholderStyle
	ti.Cursor.Style = styles.TextInputCursorStyle

	return Model{
		newFIleInput:     ti,
		fileInputVisible: false,
		noteTextArea:     ta,
		noteList:         noteList,
	}
}
