package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"

	"github.com/RudraPratapDev/NoteX/internal/config"
	"github.com/RudraPratapDev/NoteX/internal/models"
)

func ListFiles() []list.Item {
	items := make([]list.Item, 0)
	entries, err := os.ReadDir(config.VaultDir)
	if err != nil {
		log.Fatal("error reading notes", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if entry.Name() == ".DS_Store" {
				continue
			}
			modTime := info.ModTime().Format("15:04 2005-01-02")
			items = append(items, models.NewItem(
				entry.Name(),
				fmt.Sprintf("Modified at : %s", modTime),
			))
		}
	}
	return items
}
