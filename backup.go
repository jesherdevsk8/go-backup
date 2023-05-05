package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Define the source directory to be backed up
	srcDir := "/home/jesherpinkman/navebild/"

	// Define the backup file name and location
	backupFileName := "backup-go.zip"
	backupDir := "/tmp/"

	// Create a new backup file
	backupFile, err := os.Create(filepath.Join(backupDir, backupFileName))
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return
	}
	defer backupFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(backupFile)
	defer zipWriter.Close()

	// Walk through the source directory and add files to the zip file
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Get the relative path of the file
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// Create a new file in the zip archive
		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// Copy the file to the zip archive
		_, err = io.Copy(zipFile, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error walking source directory:", err)
		return
	}

	fmt.Println("Backup created successfully!")
}
