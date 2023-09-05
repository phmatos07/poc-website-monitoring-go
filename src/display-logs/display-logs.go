package displaylogs

import (
	"fmt"
	"io/ioutil"

	"github.com/phmatos07/poc-website-monitoring-go/errorLog"
)

func ToView() {
	file, err := ioutil.ReadFile("log-site.txt")

	if err != nil {
		errorLog.Log(err)
	}

	fmt.Println("")
	fmt.Println("LOGS:")
	fmt.Println(string(file))
}
