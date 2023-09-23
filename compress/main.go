package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Input file path (the file you want to zip)
	inputFilePath := "C:/Users/Akash Heliware/Documents/Go_tutoral/compress/input_file/Input.txt"

	// Output zip file path (where you want to save the compressed file)
	outputZipPath := "C:/Users/Akash Heliware/Documents/Go_tutoral/compress/ouput_file/zipfile.zip" //compress/ouput_file

	// Create or truncate the output zip file
	outputZipFile, err := os.Create(outputZipPath)
	if err != nil {
		fmt.Println("Error creating output zip file:", err)
		return
	}
	defer outputZipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(outputZipFile)
	defer zipWriter.Close()

	// Open the input file for reading
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Get the filename from the input path
	inputFileName := filepath.Base(inputFilePath)

	// Create a file header for the zip file
	fileHeader := &zip.FileHeader{
		Name:   inputFileName,
		Method: zip.Store, // Compression method (zip.Store for no compression)
	}

	// Create a writer for the zip file
	zipFileWriter, err := zipWriter.CreateHeader(fileHeader)
	if err != nil {
		fmt.Println("Error creating zip file entry:", err)
		return
	}

	// Copy the contents of the input file to the zip file
	_, err = io.Copy(zipFileWriter, inputFile)
	if err != nil {
		fmt.Println("Error copying file contents to zip:", err)
		return
	}

	fmt.Printf("File '%s' has been compressed and saved to '%s'\n", inputFileName, outputZipPath)
}
