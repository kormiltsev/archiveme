package compressers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Compresso find file or folder by address provided, tar it, and gzip it. Results in buf. Returns error.
func Compresso(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// is file a folder?
	fi, err := os.Stat(src)
	if err != nil {
		log.Println("Compresso stat error:", err)
		return err
	}
	mode := fi.Mode()
	if mode.IsRegular() {
		// get header
		header, err := tar.FileInfoHeader(fi, src)
		if err != nil {
			log.Println("FileInfoHeader error:", err)
			return err
		}
		// write header
		if err := tw.WriteHeader(header); err != nil {
			log.Println("WriteHeader error:", err)
			return err
		}

		// get content
		data, err := os.Open(src)
		if err != nil {
			log.Println("Open error:", err)
			return err
		}
		if _, err := io.Copy(tw, data); err != nil {
			log.Println("Copy error:", err)
			return err
		}
	} else if mode.IsDir() { // folder

		// walk through every file in the folder
		filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
			// generate tar header
			header, err := tar.FileInfoHeader(fi, file)
			if err != nil {
				log.Println("WriteHeader err:", err)
				return err
			}

			// must provide real name
			header.Name = filepath.ToSlash(file)

			// write header
			if err := tw.WriteHeader(header); err != nil {
				log.Println("WriteHeader err:", err)
				return err
			}
			// if not a dir, write file content
			if !fi.IsDir() {
				data, err := os.Open(file)
				if err != nil {
					log.Println("Walk Open err:", err)
					return err
				}
				if _, err := io.Copy(tw, data); err != nil {
					log.Println("Copy err:", err)
					return err
				}
			}
			return nil
		})
	} else {
		return fmt.Errorf("error: file type not supported")
	}

	// produce tar
	if err := tw.Close(); err != nil {
		log.Println("tw.Close err:", err)
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		log.Println("zr.Close err:", err)
		return err
	}
	//
	return nil
}
