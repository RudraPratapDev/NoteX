package config

import (
	"fmt"
	"log"
	"os"
)

var VaultDir string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
	}
	VaultDir = fmt.Sprintf("%s/.notex", homeDir)
}
