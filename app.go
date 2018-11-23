package main

import (
	"time"
	"github.com/go-redis/redis"
	"math/rand"
)

type App struct {
	id string
	isGenerator bool
	client *redis.Client
}


func launchApp() *App {
	app := new(App)
	app.id = RandStringBytesMask(8)
	app.isGenerator = false
	app.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return app
}


const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}


//Controller
func (app *App) Controller () {
	if app.isGenerator {
		app.sendMessage()
		time.Sleep(time.Millisecond * 500)
	}else{
		app.verifyMessage()
		time.Sleep(time.Millisecond * 4)
	}
}
