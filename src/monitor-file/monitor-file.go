package monitorFile

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/phmatos07/poc-website-monitoring-go/errorLog"
	"github.com/phmatos07/poc-website-monitoring-go/monitoring"
)

func ToMonitor() {

	sites := readFile()

	for i := 0; i < len(sites); i++ {
		resp, err := http.Get(sites[i])

		if err != nil {
			errorLog.Log(err)
		}
		monitoring.ToView(sites[i], resp.StatusCode)
	}
}

func readFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		errorLog.Log(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}
