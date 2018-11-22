package main

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

func (app *App) checkStatus () {
	generatorID, err := app.client.Get("generatorID").Result()
	if err == redis.Nil {
		app.setStatus()
	} else if err != nil {
		fmt.Println("ERROR occurred while CHECKING the generator status!: ", err)
	} else if generatorID == app.id {
		app.updateStatus()
	} else {
		app.isGenerator = false
	}
}

func (app *App) setStatus() {
	reply := app.client.SetNX("generatorID", app.id, time.Millisecond * 1000)
	if err := reply.Err(); err != nil {
		fmt.Println("ERROR occurred while STARTING the generator!", err)
		app.isGenerator = false
	} else if reply.Val() == false {
		fmt.Println("Attempt to start generator FAILED! Generator has been started already")
		app.isGenerator = false
	} else {
		fmt.Println("Generator was started successfully!")
		app.isGenerator = true
	}
}

func (app *App) updateStatus() {
	reply := app.client.Expire("generatorID", time.Millisecond * 1000)
	if err := reply.Err(); err != nil {
		fmt.Println("ERROR occurred while UPDATING the generator status!")
		fmt.Println("The generator operation is terminated.")
		app.isGenerator = false
	}
}
