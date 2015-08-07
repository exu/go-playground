package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	filePath := "test.jpg"
	if reader, err := os.Open(filePath); err == nil {
		defer reader.Close()

		im, _, err := image.DecodeConfig(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", filePath, err)
			return
		}
		fmt.Printf("%s %d %d\n", filePath, im.Width, im.Height)
	} else {
		fmt.Println("Impossible to open the file")
	}

}
