package logs

import (
	"fmt"
	"log"
	"os"
)

func WriteToLogFile(msg string) {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
	fmt.Println(msg)
}
