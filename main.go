package main

import (
	"fmt"
	"go-tube/internal/app"
	"go-tube/internal/lib"
	"os"
)

func main() {
	if lib.IsDev {
		lib.ClearConsole()
	}

	program := app.StartApp()
	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program")
		os.Exit(1)
	}
}
