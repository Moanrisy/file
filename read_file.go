package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Tidak dapat membuka file")
	}

	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	for fileScan.Scan() {
		fmt.Println(fileScan.Text())
	}
}
