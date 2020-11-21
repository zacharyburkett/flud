package icmpflut

import (
	"image"
	"log"
	"sync"

	"github.com/mehrdadrad/ping"
)

type Renderer struct {
	ips   <-chan string
	index int

	options Options
}

func NewRenderer(img image.Image, options Options) *Renderer {
	return &Renderer{
		ips:     imageToIPs(img, options),
		options: options,
	}
}

func (r *Renderer) Draw() error {
	var wg sync.WaitGroup
	wg.Add(r.options.Workers)

	for i := 0; i < r.options.Workers; i++ {
		go func() {
			defer wg.Done()

			for {
				ip, ok := <-r.ips
				if !ok {
					return
				}

				p, err := ping.New(ip)
				if err != nil {
					log.Println(err)
					continue
				}

				ping := func() bool {
					p.SetCount(1)
					p.SetTimeout(r.options.Timeout)

					resp, err := p.Run()
					if err != nil {
						log.Println(err)
						return false
					}

					r := <-resp
					if r.Err == nil {
						log.Println(r)
						return true
					}

					return false
				}

				for i := 0; i < r.options.Count; i++ {
					if ping() {
						break
					}
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
