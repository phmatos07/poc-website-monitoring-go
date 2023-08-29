package main

import (
	"fmt"

	"github.com/phmatos07/poc-website-monitoring-go/command"
)

func main() {
	var data = command.Get()
	fmt.Println(data)
}
