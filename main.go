package main

import (
	"flag"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func listenForQuit() {
	println("Press q at anytime to quit")
	if robotgo.AddEvent("q") {
		println("Quiting...")
		os.Exit(0)
	}
}

func main() {
	delay := flag.Int("d", 2, "delay in seconds")
	flag.Parse()
	slumber := time.Duration(*delay) * time.Second
	go listenForQuit()

	for true {
		time.Sleep(slumber)
		robotgo.MouseClick("right", true)
	}
}
