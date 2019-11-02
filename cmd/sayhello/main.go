package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	times = kingpin.Flag("times", "How many times to say hello. Env: TIMES").
		Default("1").Short('t').Envar("TIMES").Uint32()
	name = kingpin.Arg("name", "Name of user.").Required().String()
)

func main() {
	kingpin.CommandLine.Name = "sayhello"
	kingpin.Parse()

	for i := *times; i > 0; i-- {
		fmt.Printf("Hello %s!\n", *name)
	}
}
