package icmpflut

import (
	"image"
	"log"
	"net"

	"github.com/mehrdadrad/ping"
)

type Renderer struct {
	ips   []net.IP
	index int
}

func NewRenderer(img image.Image) *Renderer {
	return &Renderer{
		ips: imageToIPs(img),
	}
}

func (r *Renderer) Draw() error {
	ip := r.nextIP().To16().String()
	p, err := ping.New(ip)
	if err != nil {
		return err
	}

	p.SetCount(1)
	p.SetTimeout("10s")

	resp, err := p.Run()
	if err != nil {
		return err
	}
	log.Println(<-resp)

	return nil
}

func (r *Renderer) nextIP() net.IP {
	ip := r.ips[r.index]

	r.index++
	if r.index >= len(r.ips) {
		r.index = 0
	}

	return ip
}
