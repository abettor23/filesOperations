package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл, возможно он отсутствует:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Size() > 0 {
		buf := make([]byte, fileInfo.Size())
		_, err = io.ReadFull(file, buf)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf))
	} else {
		fmt.Println("Файл пустой.")
	}
}
