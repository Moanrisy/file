package service

import (
	"fmt"
	"os"
)

func AppendText(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	resultOfText := fmt.Sprintf("\n%s", data)

	if _, err := file.WriteString(resultOfText); err != nil {
		fmt.Println(err)
	}
}
