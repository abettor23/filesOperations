package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.Create("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Введите свое сообщение ниже. Для выхода введите exit.")

	for i := 1; ; i++ {

		var buf bytes.Buffer

		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}

		if fileInfo.Size() > 0 {
			buf.WriteString(fmt.Sprintf("\n"))
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(strings.TrimSpace(input)) == "exit" {
			return
		} else if strings.ReplaceAll(input, " ", "") == "" {
			continue
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		buf.WriteString(fmt.Sprintf("%d %s %s", i, currentTime, input))
		_, err = file.Write(buf.Bytes())
		if err != nil {
			panic(err)
		}
	}
}
