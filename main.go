package main

import (
	"time"
)

func main() {

	app := launchApp()

	go func() {
		for {
			app.checkStatus()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for {
		app.Controller()
	}

}
