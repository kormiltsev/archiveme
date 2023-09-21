package app

import (
	"flag"
	"fmt"
)

// parseflag parses flags.
func (cod *Coder) parseflag() error {

	// Client settings
	flag.StringVar(&cod.SourceName, "f", "", "File or Folder to encode OR file.archiveme to decode")

	flag.StringVar(&cod.Password, "p", cod.Password, fmt.Sprintln("your password"))

	flag.StringVar(&cod.Result, "d", cod.Result, fmt.Sprintln("result file"))

	flag.StringVar(&cod.FileType, "type", cod.FileType, fmt.Sprintln("Type of encoded files"))

	flag.Parse()

	if cod.SourceName == "" {
		if len(flag.Args()) == 0 {
			// fmt.Print(about())
			return fmt.Errorf("at least one argument required (./fileToArchive)")
		}
		cod.SourceName = flag.Args()[0]
	}
	return nil
}
