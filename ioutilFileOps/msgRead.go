package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "messages.txt"
	fileInfo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла, возможно он отсутствует:", err)
		return
	} else if len(fileInfo) == 0 {
		fmt.Println("Файл пустой.")
	} else {
		fmt.Println(string(fileInfo))
	}
}
