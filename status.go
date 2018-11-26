package main

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

//Check app status at Redis database:
//If there is no "generatorID" key the app will try to become generator;
//If the value on "generatorID" key is equal to app.id then the app is generator already and status wil be updated;
//In any other case the app is not a generator.
func (app *App) checkStatus () (bool, error) {
	generatorID, err := app.client.Get("generatorID").Result()

	switch {
	case err == redis.Nil :
		return statusFunc(app.setStatus)

	case err != nil :
		fmt.Println("ERROR occurred while CHECKING the generator status!: ", err)
		return false, err

	case generatorID == app.id :
		return statusFunc(app.updateStatus)

	default :
		app.isGenerator = false
		return true, err
	}
}

func statusFunc(f func() (bool, error)) (bool, error) {
	if ok, err := f(); ok {
		return true, err
	} else {
		return false, err
	}
}

//Try to set app.id as generatorID to Redis database only if the "generatorID" key does not exist.
//If redis cmd is successful then the app has become a generator;
//If "generatorID" key exists already then the app is not a generator.
func (app *App) setStatus() (bool, error) {
	reply := app.client.SetNX("generatorID", app.id, time.Millisecond * 1000)
	if err := reply.Err(); err != nil {
		app.isGenerator = false
		return false, newError("/n ERROR occurred while STARTING THE GENERATOR! /n"/*, err*/)

	} else if reply.Val() == false {
		app.isGenerator = false
		return true, newError("/n Attempt to start generator FAILED! Generator has been started already /n")

	} else {
		fmt.Println("/n Generator has been started successfully! /n")
		app.isGenerator = true
		return true, nil
	}
}

//Update the app status by extending the existence of the "generatorID" key
func (app *App) updateStatus() (bool, error) {
	reply := app.client.Expire("generatorID", time.Millisecond * 1000)
	if err := reply.Err(); err != nil {
		//fmt.Println("/n The generator operation is terminated. /n")
		app.isGenerator = false
		return false, newError("ERROR occurred while UPDATING THE GENERATOR STATUS! /n" /*+ err*/)
	} else {
		return true, nil
	}
}
