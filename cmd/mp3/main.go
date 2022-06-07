package main

import (
	"fmt"
	"tools/media"
)

func main() {
	size, err := media.Download("https://test.mp3", "test.mp3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("size = %f M\n", float64(size/1024.0/1024))
}
