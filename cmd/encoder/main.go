package main

import (
	"fmt"
	"os"

	"github.com/kormiltsev/archiveme/internal/app"
)

// errf operates errors and return Exit status 1.
func errf(err error) {
	fmt.Println("Error:", err)
	fmt.Print(app.About())
	os.Exit(1)
}

func main() {

	// return new coder structure with settings
	coder, err := app.New()
	if err != nil {
		errf(err)
	}

	// run arviver
	err = coder.Do()
	if err != nil {
		errf(err)
	}

	fmt.Printf("Result is here %s\n", coder.Result)
}
