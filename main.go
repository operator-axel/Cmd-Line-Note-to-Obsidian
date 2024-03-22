package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What should the name of the file be? ")
	fileName, _ := reader.ReadString('\n')
	if !strings.HasSuffix(fileName, ".md") {
		fileName += ".md"
	}

	filePath := fmt.Sprintf("/home/ch40s/ch40s_linux/Slipbox/%s", fileName)

	fmt.Print("Note: ")
	content, _ := reader.ReadString('\n')

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

}
