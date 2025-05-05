package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var supportedExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".mp4", ".mov", ".avi", ".mkv", ".webm"}

func isSupportedFile(ext string) bool {
	ext = strings.ToLower(ext)
	for _, e := range supportedExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func fileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func moveFile(src, dst string) error {
	dstPath := filepath.Join(dst, filepath.Base(src))
	return os.Rename(src, dstPath)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the full path to your folder: ")
	sourceDir, _ := reader.ReadString('\n')
	sourceDir = strings.TrimSpace(sourceDir)

	// Ensure the input directory exists
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Println("Error: Source directory does not exist.")
		return
	}

	destDir := filepath.Join(sourceDir, "filtered_files")

	// Create the destination folder if it doesn't exist
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			fmt.Println("Failed to create filtered_files folder:", err)
			return
		}
	}

	seenHashes := make(map[string]bool)
	counter := 0

	err := filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Walk error:", err)
			return nil
		}

		if d.IsDir() || !isSupportedFile(filepath.Ext(d.Name())) || filepath.Dir(path) == destDir {
			return nil
		}

		hash, err := fileHash(path)
		if err != nil {
			fmt.Println("Error hashing file:", path, err)
			return nil
		}

		if seenHashes[hash] {
			fmt.Println("Duplicate skipped:", d.Name())
			return nil
		}

		err = moveFile(path, destDir)
		if err != nil {
			fmt.Println("Error moving file:", path, err)
			return nil
		}

		counter++

		seenHashes[hash] = true
		fmt.Printf("Moved %d to filtered_files: %s \n", counter, d.Name())
		return nil
	})

	if err != nil {
		fmt.Println("Error during processing:", err)
	}
}
