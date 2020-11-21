package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: flud <subcommand> <args>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "icmp":
		icmpCmd.Parse(os.Args[2:])
		execICMP()

	default:
		fmt.Printf("%s is not a valid subcommand\n", os.Args[1])
		os.Exit(1)
	}
}
