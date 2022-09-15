package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("defer.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			os.Stdout.Write([]byte("\n"))
			break
		}
	}
}
