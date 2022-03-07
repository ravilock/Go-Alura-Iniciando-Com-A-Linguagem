package main

import (
	"fmt"
	"os"

	"cursoAllura/src/logger"
	"cursoAllura/src/siteInspector"
)

func main() {
	for {
		showMenu()
		command := getCommand()

		switch command {
		case 1:
			siteInspector.StartMonitoring()
		case 2:
			logger.LoggerRead()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido.")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}
