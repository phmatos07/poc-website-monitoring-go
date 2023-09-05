package monitoring

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/phmatos07/poc-website-monitoring-go/errorLog"
	timebr "github.com/phmatos07/poc-website-monitoring-go/time-br"
)

func ToMonitor() {

	var site, command string

	fmt.Println("Digite o endereço referente ao site")
	fmt.Scan(&site)

	for {
		resp, err := http.Get(site)

		if err != nil {
			errorLog.Log(err)
		}

		ToView(site, resp.StatusCode)

		questions()
		fmt.Scan(&command)

		if command != "y" {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		}
	}
}

func ToView(site string, statusCode int) {

	var mensage string

	if statusCode == 200 {
		mensage = timebr.ToView() + " - Site: " + site + " foi carregado com sucesso! | Status Code: " + strconv.FormatInt(int64(statusCode), 10) + "\n"
	} else {
		mensage = timebr.ToView() + "Site: " + site + " está com problemas! | Status Code: " + strconv.FormatInt(int64(statusCode), 10) + "\n"
	}

	fmt.Println(mensage)
	errorLog.RegisterLogSite(mensage)
}

func questions() {
	fmt.Println("")
	fmt.Println("Desejar repetir o processo?")
	fmt.Println("Obs.: Digite 'y' para repetir ou qualquer coisa para cancelar")
	fmt.Println("")
}
