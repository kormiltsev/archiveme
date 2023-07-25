package app

import (
	"flag"
	"fmt"
)

func (cod *Coder) parseflag() error {

	// Client settings
	flag.StringVar(&cod.SourceName, "f", "", "File or Folder to encode OR file.nau to decode")

	flag.StringVar(&cod.Password, "p", cod.Password, fmt.Sprintf("your password (default %s )", cod.Password))

	flag.StringVar(&cod.Result, "d", cod.Result, fmt.Sprintf("Directory for results (default %s )", cod.Result))

	flag.StringVar(&cod.FileType, "type", cod.FileType, fmt.Sprintf("Type of encoded files (default %s )", cod.FileType))

	flag.Parse()

	if cod.SourceName == "" {
		if len(flag.Args()[0]) == 0 {
			// fmt.Print(about())
			return fmt.Errorf("Zero arguments")
		}
		cod.SourceName = flag.Args()[0]
	}
	return nil
}
