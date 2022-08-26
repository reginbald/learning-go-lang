package main

import (
	"os"

	"reginbald.com/internal/images"
	"reginbald.com/internal/linkedList"
	"reginbald.com/internal/loopsAndFunctions"
	"reginbald.com/internal/rot13Reader"
	"reginbald.com/internal/tree"
	"reginbald.com/internal/webCrawler"
)

func main() {
	switch os.Args[1] {
	case "loopsAndFunctions":
		loopsAndFunctions.Run()
	case "rot13Reader":
		rot13Reader.Run()
	case "images":
		images.Run()
	case "linkedList":
		linkedList.Run()
	case "tree":
		tree.Run()
	case "webCrawler":
		webCrawler.Run()
	}
}
