package errorLog

import (
	"fmt"
	"os"
)

const mensage = "Infelizmente ocorreu um erro!"

func Log(err error) {
	fmt.Println(mensage)
	fmt.Println("Erro: ", err)
	registerError(err.Error())
	os.Exit(0)
}

func RegisterLogSite(mensage string) {

	file, err := os.OpenFile("log-site.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		Log(err)
	}
	file.WriteString(mensage)
	file.Close()
}

func registerError(erro string) {

	file, err := os.OpenFile("error-log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(mensage)
		fmt.Println("Erro: ", err)
		file.WriteString("Erro: " + err.Error() + "\n")
		os.Exit(0)
	}
	file.WriteString("Erro: " + erro + "\n")
	file.Close()
}
