package errorLog

import (
	"fmt"
	"os"
)

func Log(err error) {
	fmt.Println("Infelizmente ocorreu um erro!")
	fmt.Println("Erro: ", err)
	os.Exit(0)
}
