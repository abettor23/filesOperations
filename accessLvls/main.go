package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	//создаю файл, пишу любую строку "чтоб было", присваиваю права доступа
	fileName := "readonly.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Не удалось создать файл.", err)
		return
	}

	var buf bytes.Buffer
	buf.WriteString("New file readonly\n")
	_, err = file.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	if err := file.Chmod(0444); err != nil {
		fmt.Println("Не удалось изменить права доступа.", err)
		return
	}

	//после всего закрываю файл
	if err := file.Close(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Файл %s был успешно создан.\n", fileName)

	//открываю файл для взаимодействия
	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Не удалось открыть файл, возможно он отсутствует:", err)
		return
	}
	defer file.Close()

	//пробую записать строку для проверки прав доступа
	buf.WriteString("Check access level.\n")
	_, err = file.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Запись в файл невозможна. Только для чтения.", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	//считываю информацию по содержимому файла, в том числе убеждаюсь, что новая строка не записана
	if fileInfo.Size() > 0 {
		buf := make([]byte, fileInfo.Size())
		_, err = io.ReadFull(file, buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf))
	} else {
		fmt.Println("Файл пустой.")
	}
}
