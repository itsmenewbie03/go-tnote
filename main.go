package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func generateNoteFilename() string {
	now := time.Now()
	return fmt.Sprintf("INTERNSHIP_NOTE_%s.md", now.Format("01-02-2006"))
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func createFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()
	fmt.Printf("File '%s' created successfully.\n", filePath)
	return nil
}

func main() {
	obsidianVaultPath := "/home/itsmenewbie03/obsidian-vault/"
	todaysNote := generateNoteFilename()
	todaysNotePath := fmt.Sprintf("%s%s", obsidianVaultPath, todaysNote)
	if pathExists(todaysNotePath) {
		dir := filepath.Dir(todaysNotePath)
		file := filepath.Base(todaysNotePath)
		fmt.Println(dir, file)
		cmd := exec.Command("nvim", file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Dir = dir
		cmd.Run()

	} else {
		err := createFile(todaysNotePath)
		if err != nil {
			fmt.Println("Error:", err)
		}
		cmd := exec.Command("nvim", todaysNotePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
