package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {

	fileName := ("messages.txt")
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Введите свое сообщение ниже. Для выхода введите exit.")

	for i := 1; ; i++ {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(strings.TrimSpace(input)) == "exit" {
			return
		} else if strings.ReplaceAll(input, " ", "") == "" {
			continue
		}

		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}

		currentStr, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		} else if fileInfo.Size() > 0 {
			currentStr = append(currentStr, []byte(fmt.Sprintf("\n"))...)
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf("%d %s %s", i, currentTime, input)

		newStr := []byte(message)
		mainStr := append(currentStr, newStr...)

		if err := ioutil.WriteFile(fileName, mainStr, 0644); err != nil {
			panic(err)
		}
	}
}
