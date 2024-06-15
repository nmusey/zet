package main

import (
	"fmt"
	"os"

	"github.com/nmusey/zet/pkg/app"
)

func main() {
	app := app.BuildApp()
	if err := app.Run(os.Args); err != nil {
        fmt.Printf("%v\n", err)
	}
}
