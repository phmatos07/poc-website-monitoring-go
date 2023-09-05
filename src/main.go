package main

import (
	"fmt"
	"os"

	"github.com/phmatos07/poc-website-monitoring-go/command"
	monitorFile "github.com/phmatos07/poc-website-monitoring-go/monitor-file"
	"github.com/phmatos07/poc-website-monitoring-go/monitoring"
)

func main() {

	_command := command.Get()

	if _command == command.TypesCommands[1] {
		monitoring.ToMonitor()
		return
	}

	if _command == command.TypesCommands[2] {
		monitorFile.ToMonitor()
		return
	}

	if _command == command.TypesCommands[3] {
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	}
}
