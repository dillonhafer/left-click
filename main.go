package main

import (
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func listenForQuit() {
	println("Press q at anytime to quit")
	if robotgo.AddEvent("q") {
		os.Exit(0)
	}
}

func main() {
	go listenForQuit()

	for true {
		time.Sleep(2 * time.Second)
		robotgo.MouseClick("left", true)
		println("left clicked")
	}
}
