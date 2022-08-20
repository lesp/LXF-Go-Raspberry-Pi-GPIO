package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"time"
)

var (
	red = rpio.Pin(17)
	yellow = rpio.Pin(27)
	green = rpio.Pin(22)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	red.Output()
	yellow.Output()
	green.Output()

	for true {
		red.Toggle()
		time.Sleep(time.Second / 5)
		yellow.Toggle()
                time.Sleep(time.Second / 5)
		green.Toggle()
                time.Sleep(time.Second / 5)
	}
}
