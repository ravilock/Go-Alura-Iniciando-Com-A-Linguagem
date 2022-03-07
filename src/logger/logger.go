package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoggerWrite(site string, siteStatus string, statusCode int) {
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return
	}
	defer file.Close()

	logLine := fmt.Sprintln(time.Now().Format("02/01/2006 15:04:05"), "-", site, siteStatus, strconv.FormatInt(int64(statusCode), 10))
	file.WriteString(logLine)

	file.Close()
}

func LoggerRead() {
	fmt.Println("Exibindo Logs...")
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		logLine, err := reader.ReadString('\n')
		logLine = strings.TrimSpace(logLine)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro", err)
			return
		}

		fmt.Println(logLine)
	}
}
