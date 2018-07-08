package engine

import (
	"os"
	"log"
	"bufio"
)

func Script(filename string) string {	
	file, err := os.Open("./views/scripts/"+filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := make([]byte, 1000000)
	reader := bufio.NewReader(file)
	count, _ := reader.Read(buf)
	return string(buf[:count])
}