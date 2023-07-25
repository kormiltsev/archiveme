package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/kormiltsev/directoryToGzipEncrypted/internal/compressers"
	"github.com/kormiltsev/directoryToGzipEncrypted/internal/encoders"
)

type Coder struct {
	Password   string
	SourceName string
	Result     string
	FileType   string
}

func New() (*Coder, error) {
	coder := new(Coder)
	coder.setDefault()

	err := coder.parseflag()
	if err != nil {
		return nil, err
	}

	return coder, nil
}

func (cod *Coder) Do() error {
	return cod.switcher()
}

func (cod *Coder) switcher() error {
	ext := strings.ToLower(filepath.Ext(cod.SourceName))
	if ext == cod.FileType {
		return cod.decod()
	}
	return cod.encod()
}

func (cod *Coder) encod() error {

	// compress file or folder
	var buf bytes.Buffer
	if err := compressers.Compresso(cod.SourceName, &buf); err != nil {
		log.Println("encod err:", err)
		return err
	}

	// encrypto
	massa, err := encoders.Encrypto(buf.Bytes(), []byte(cod.Password))
	if err != nil {
		log.Println("Encrypto err:", err)
		return err
	}

	// write file to disk
	_, outFile := filepath.Split(cod.SourceName)
	outFile = outFile + cod.FileType
	if err = ioutil.WriteFile(outFile, massa, 0600); err != nil {
		log.Println(" encod WriteFile err:", err)
		return err
	}
	return nil
}

func (cod *Coder) decod() error {
	// open file
	massa, err := ioutil.ReadFile(cod.SourceName)
	if err != nil {
		return err
	}

	// decrypto
	archive, err := encoders.Decrypto(massa, []byte(cod.Password))
	if err != nil {
		return err
	}

	// decompress it
	buf2 := bytes.NewReader(archive)
	if err := compressers.Decompresso(buf2, cod.Result); err != nil {
		return err
	}

	return nil
}
