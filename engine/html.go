package engine

import (
	"os"
	"log"
	"bufio"
)

func HtmlIndex() string {	
	file, err := os.Open("./views/html/index.html")
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := make([]byte, 10000)
	reader := bufio.NewReader(file)
	reader.Read(buf)
	return string(buf)
}