package icmpflut

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func imageToIPs(img image.Image, options Options) <-chan string {
	ips := make(chan string, options.Workers)

	var template = options.Prefix + "%04d:%04d:%02x%02x:%02x"

	go func() {
		var buffer []string

		for x := 0; x < img.Bounds().Dx(); x++ {
			for y := 0; y < img.Bounds().Dy(); y++ {
				r, g, b, a := img.At(x, y).RGBA()
				if a == 0 {
					continue
				}

				ip := fmt.Sprintf(template, x+options.X, y+options.Y, r/257, g/257, b/257)
				buffer = append(buffer, ip)
			}
		}

		rand.Shuffle(len(buffer), func(i, j int) {
			buffer[i], buffer[j] = buffer[j], buffer[i]
		})

		for {
			for _, ip := range buffer {
				ips <- ip
			}
		}
	}()

	return ips
}
