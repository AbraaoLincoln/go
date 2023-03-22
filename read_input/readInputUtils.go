package main

import (
	"bufio"
	"log"
	"os"
)

type InputProcessor interface {
	process(value string)
}

func readInput(commandLineArgs []string, inputProcessor InputProcessor) {
	if len(commandLineArgs) > 1 {
		readFromFile(commandLineArgs[1], inputProcessor)
	} else {
		readFromStdin(inputProcessor)
	}
}

func readFromStdin(inputProcessor InputProcessor) {
	log.Println("Reading from stdin")

	read(os.Stdin, inputProcessor)

	log.Println("Finished reading")
}

func readFromFile(fileName string, inputProcessor InputProcessor) {
	log.Println("Reading from file:", fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	read(file, inputProcessor)

	log.Println("Finished reading")
}

func read(file *os.File, inputProcessor InputProcessor) {
	log.Println("Reading input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputProcessor.process(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
