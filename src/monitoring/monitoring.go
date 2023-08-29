package monitoring

import (
	"fmt"
	"net/http"
	"os"
)

func ToMonitor() {

	var site, command string

	fmt.Println("Digite o endereço referente ao site")
	fmt.Scan(&site)

	for {
		resp, _ := http.Get(site)
		toView(site, resp.StatusCode)

		questions()
		fmt.Scan(&command)

		if command != "y" {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		}
	}
}

func toView(site string, statusCode int) {
	if statusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", statusCode)
	}
}

func questions() {
	fmt.Println("")
	fmt.Println("Desejar repetir o processo?")
	fmt.Println("Obs.: Digite 'y' para repetir ou qualquer coisa para cancelar")
	fmt.Println("")
}
