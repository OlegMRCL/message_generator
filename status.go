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
		return false, err
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
		fmt.Println("ERROR occurred while STARTING the generator!", err)
		app.isGenerator = false
		return false, err
	} else if reply.Val() == false {
		fmt.Println("Attempt to start generator FAILED! Generator has been started already")
		app.isGenerator = false
		return false, err
	} else {
		fmt.Println("Generator was started successfully!")
		app.isGenerator = true
		return true, err
	}
}

//Update the app status by extending the existence of the "generatorID" key
func (app *App) updateStatus() (bool, error) {
	reply := app.client.Expire("generatorID", time.Millisecond * 1000)
	if err := reply.Err(); err != nil {
		fmt.Println("ERROR occurred while UPDATING the generator status!")
		fmt.Println("The generator operation is terminated.")
		app.isGenerator = false
		return true, err
	} else {
		return true, err
	}
}
