package compressers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// check for path correct
func validRelPath(p string) bool {
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimPrefix(p, "../")

	if p == "" || strings.Contains(p, `\`) || strings.Contains(p, "../") {
		return false
	}

	return true
}

// Decompresso find file arcived by address provided, decode, unzip and untar it to the result folder. Returns error.
func Decompresso(src io.Reader, dst string) error {
	// ungzip
	zr, err := gzip.NewReader(src)
	if err != nil {
		return fmt.Errorf("can't decompress file, wrong password? error: %s", err)
	}
	// untar
	tr := tar.NewReader(zr)

	// uncompress each element
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		target := header.Name

		// validate name against path traversal
		if !validRelPath(header.Name) {
			return fmt.Errorf("tar contained invalid name error %q", target)
		}

		// add dst + re-format slashes according to system
		target = filepath.Join(dst, header.Name)

		// check the type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it (with 0755 permission)
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					fmt.Println("create directory error")
					return err
				}
			}
		// if it's a file create it (with same permission)
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}

			fileToWrite.Close()
		}
	}

	return nil
}
