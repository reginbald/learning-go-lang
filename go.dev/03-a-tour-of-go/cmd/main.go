package main

import (
	"os"

	"reginbald.com/internal/images"
	"reginbald.com/internal/linkedList"
	"reginbald.com/internal/rot13Reader"
)

func main() {
	switch os.Args[1] {
	case "0":
		rot13Reader.Run()
	case "1":
		images.Run()
	case "2":
		linkedList.Run()
	}

}
