package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"mars-rover/pkg/message"
	"mars-rover/pkg/positioning"
	"mars-rover/pkg/storage"
)

func main() {
	filePtr := flag.String("fPath", "sample.txt", "message path to read from")
	flag.Parse()
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	repository := storage.NewRepository()
	newRobot := positioning.NewRobot()

	interactor := positioning.NewInteractor(repository, newRobot)
	handler := message.NewHandler(interactor)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := handler.Processor(scanner.Text()); err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
