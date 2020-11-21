package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/zacharyburkett/flud/icmpflut"
)

var (
	icmpCmd = flag.NewFlagSet("icmp", flag.ExitOnError)

	xFlag = icmpCmd.Int("x", 0, "X position")
	yFlag = icmpCmd.Int("y", 0, "Y position")

	prefix  = icmpCmd.String("prefix", "2001:19f0:5001:36dc:", "IPv6 prefix")
	workers = icmpCmd.Int("w", 5, "Worker count")
	timeout = icmpCmd.String("t", "500ms", "Ping timeout")
	count   = icmpCmd.Int("c", 100, "Ping attempt count")
)

func execICMP() {
	if len(icmpCmd.Args()) != 1 {
		fmt.Println("Usage: flud icmp <args> <filename>")
		os.Exit(1)
	}

	f, err := os.Open(icmpCmd.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	renderer := icmpflut.NewRenderer(img, icmpflut.Options{
		Prefix:  *prefix,
		Timeout: *timeout,
		Workers: *workers,
		Count:   *count,
		X:       *xFlag,
		Y:       *yFlag,
	})
	renderer.Draw()
}
