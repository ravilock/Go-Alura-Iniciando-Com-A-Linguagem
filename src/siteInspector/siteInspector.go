package siteInspector

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"cursoAllura/src/logger"
)

const numberOfWebsiteChecks = 5
const delayBetweenWebsiteChecks = 5 * time.Second

func StartMonitoring() {
	fmt.Println("Monitorando...")
	sites := extractSitesFromFile("./sites.txt")

	for i := 0; i < numberOfWebsiteChecks; i++ {
		for _, site := range sites {
			testWebsiteDisponibility(site)
		}
		time.Sleep(delayBetweenWebsiteChecks)
	}
}

func extractSitesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return nil
	}
	defer file.Close()

	var sites []string
	reader := bufio.NewReader(file)
	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro", err)
			return nil
		}

		sites = append(sites, site)
	}

	file.Close()

	return sites
}

func testWebsiteDisponibility(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logger.LoggerWrite(site, "online", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		logger.LoggerWrite(site, "offline", resp.StatusCode)
	}
}
