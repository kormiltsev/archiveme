package compressers

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCompresso(t *testing.T) {
	// Create a temporary file
	tempFile := createTempFile()
	defer os.RemoveAll(tempFile) // Clean up after the test is done

	// Create a buffer to hold the compressed data
	var compressedData bytes.Buffer

	// Call the Compresso function to compress the files
	err := Compresso(tempFile, &compressedData)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check that the compressed data is not empty
	if compressedData.Len() == 0 {
		t.Errorf("Compressed data is empty")
	}

	// Verify that the compressed data can be successfully decompressed
	decompressedData, err := decompressData(&compressedData)
	if err != nil {
		t.Fatalf("Unexpected error during decompression: %v", err)
	}

	// Compare the decompressed data with the original data
	originalData := getOriginalData(tempFile)
	if !bytes.Equal(decompressedData, originalData) {
		t.Errorf("Decompressed data does not match the original data")
	}
}

func createTempFile() string {
	// Create file
	file := filepath.Join(".", "file1.txt")

	data := []byte("This is file 1.")

	ioutil.WriteFile(file, data, 0644)

	return file
}

func decompressData(compressedData *bytes.Buffer) ([]byte, error) {
	gr, err := gzip.NewReader(compressedData)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	// Read the contents of the tar archive and store it in a buffer
	var decompressedData bytes.Buffer
	io.Copy(&decompressedData, tr)

	return decompressedData.Bytes(), nil
}

func getOriginalData(tempDir string) []byte {
	// Read the contents of the original files and concatenate them into a single buffer
	var originalData bytes.Buffer
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")

	data1, _ := ioutil.ReadFile(file1)
	data2, _ := ioutil.ReadFile(file2)

	originalData.Write(data1)
	originalData.Write(data2)

	return originalData.Bytes()
}
