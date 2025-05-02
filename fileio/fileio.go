package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text!")
	file.Close()

	stream, err := os.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(stream)

}
