package main

import (
	"image/png"
	"log"
	"os"

	"github.com/zacharyburkett/flud/icmpflut"
)

func main() {
	f, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	renderer := icmpflut.NewRenderer(img)

	for {
		if err := renderer.Draw(); err != nil {
			log.Println(err)
		}
	}
}
