package main

import (
	"fmt"
	"os"

	"github.com/phmatos07/poc-website-monitoring-go/command"
	"github.com/phmatos07/poc-website-monitoring-go/monitoring"
)

func main() {

	_command := command.Get()

	if _command == command.TypesCommands[1] {
		monitoring.ToMonitor()
		return
	}

	if _command == command.TypesCommands[2] {
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	}
}
