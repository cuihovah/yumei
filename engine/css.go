package engine

import (
	"os"
	"log"
	"bufio"
)

func Css(filename string) string {	
	file, err := os.Open("./views/styles/"+filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := make([]byte, 1000000)
	reader := bufio.NewReader(file)
	count, _ := reader.Read(buf)
	return string(buf[:count])
}