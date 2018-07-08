package engine

import (
	"os"
	"log"
	"bufio"
)

func Assets(filename string) []byte {	
	file, err := os.Open("./assets/"+filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := make([]byte, 1000000)
	reader := bufio.NewReader(file)
	count, _ := reader.Read(buf)
	return buf[:count]
}