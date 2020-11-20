package icmpflut

import (
	"fmt"
	"image"
	"net"
)

func imageToIPs(img image.Image) []net.IP {
	var ips []net.IP

	const template = "2001:41d0:2:9b23:%04d:%04d:%02x%02x:%02x"

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			ip := fmt.Sprintf(template, x, y, r/257, g/257, b/257)
			ips = append(ips, net.ParseIP(ip))
		}
	}

	return ips
}
