package main

import (
	"fmt"
	"os"

	"github.com/kormiltsev/directoryToGzipEncrypted/internal/app"
)

func errf(err error) {
	fmt.Println("Error:", err)
	app.About()
	os.Exit(1)
}

func main() {

	coder, err := app.New()
	if err != nil {
		errf(err)
	}

	err = coder.Do()
	if err != nil {
		errf(err)
	}

	fmt.Printf("Result is here %s\n", coder.Result)
}
