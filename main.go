package main

import (
	"fmt"
	"go-tube/internal/lib"
	"go-tube/internal/model"
	"os"
)

func main() {
	if lib.IsDev {
		lib.ClearConsole()
	}

	p := model.StartModel()
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program")
		os.Exit(1)
	}
}
