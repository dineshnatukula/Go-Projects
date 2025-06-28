package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	rootDir := "c://Users//dinesh.natukula//Documents//backend-services-new-repo//" // or change to your parent directory path

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if not a directory
		if !info.IsDir() {
			return nil
		}

		// Check if it's a Git repo (contains .git directory)
		gitDir := filepath.Join(path, ".git")
		if _, err := os.Stat(gitDir); os.IsNotExist(err) {
			return nil
		}

		fmt.Println("Pulling in:", path)

		// Step 1: git checkout master
		checkout := exec.Command("git", "checkout", "master")
		checkout.Dir = path
		checkout.Stdout = os.Stdout
		checkout.Stderr = os.Stderr
		if err := checkout.Run(); err != nil {
			fmt.Println("❌ Failed to checkout master in", path, ":", err)
			return nil
		}

		// Step 2: git pull
		pull := exec.Command("git", "pull")
		pull.Dir = path
		pull.Stdout = os.Stdout
		pull.Stderr = os.Stderr
		if err := pull.Run(); err != nil {
			fmt.Println("❌ Failed to pull in", path, ":", err)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directories:", err)
	}
}
