package main

import (
	"fmt"
	"os"

	"github.com/phmatos07/poc-website-monitoring-go/command"
	displaylogs "github.com/phmatos07/poc-website-monitoring-go/display-logs"
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
		displaylogs.ToView()
		return
	}

	if _command == command.TypesCommands[4] {
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	}
}
