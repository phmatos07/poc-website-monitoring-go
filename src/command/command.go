package command

import "fmt"

var TypesCommands map[int]string = map[int]string{
	0: "DESCONHECIDO",
	1: "MONITORAR_SITE",
	2: "MONITORAR_VARIOS_SITE",
	3: "EXIBIR_LOGS",
	4: "EXIT",
}

func Get() string {

	var command int
	initialQuestion := true

	for {
		questions(initialQuestion)
		fmt.Scan(&command)
		commandTransformed := toTransform(command)

		if commandTransformed == TypesCommands[0] {
			initialQuestion = false
		} else {
			return commandTransformed
		}
	}
}

func questions(initialQuestion bool) {

	if initialQuestion {
		fmt.Println("Olá, esse é o aplicativo para monitoramento de Sites, selecione a opção que atende seu cenário:")
	}
	fmt.Println("1: Monitorar um Site")
	fmt.Println("2: Monitorar vários sites (URL`s) de um arquivo.txt")
	fmt.Println("3: Exibir Logs")
	fmt.Println("4: Fechar aplicação")
}

func toTransform(command int) string {

	switch command {
	case 1:
		return TypesCommands[command]
	case 2:
		return TypesCommands[command]
	case 3:
		return TypesCommands[command]
	case 4:
		return TypesCommands[command]
	default:
		fmt.Println("Infelizmente esse comando não existe, tente novamente!")
		return TypesCommands[0]
	}
}
