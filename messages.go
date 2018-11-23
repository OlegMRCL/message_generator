package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"github.com/bxcodec/faker"
	"encoding/json"
	"math/rand"
)

type Message struct {
	GenAt string
	GenBy string
	Text  string
}

//Create new Message
func (app *App) newMessage() *Message {
	m := new (Message)
	m.GenAt = time.Now().Format("01-02-2006 15:04:05 Mon")
	m.GenBy = app.id
	m.Text = loremIpsum()
	return m
}

//Just generate random pseudo text
func loremIpsum() string {
	lorem := new (faker.Lorem)
	s := lorem.Sentence()
	return s
}

//Send the new message to queue in redis
func (app *App) sendMessage() {
	m := app.newMessage()
	JSON, _ := json.Marshal(m)
	reply := app.client.LPush("Messages", string(JSON))
	if reply.Err() != nil {
		fmt.Println("ERROR occurred while SENDING the message!")
		fmt.Println("The generator operation is terminated.")
		app.isGenerator = false
	} else {
		fmt.Println("Message has been generated: ", string(JSON))
	}
}

//Get a message from the queue and verifies it
func (app *App) verifyMessage() {
	reply := app.client.RPop("Messages")
	err := reply.Err();
	if err == redis.Nil {
		time.Sleep(time.Millisecond * 250)
	}else if err != nil {
		fmt.Println("ERROR occurred while the message VERIFIFICATION!", err)
	}else if reply.Val() != "" {
		app.verify(reply.Val())
	}
}

//Simulates message verification: identidies the message as incorrect with a 5% probability
func (app *App) verify(message string) {
	rand.Seed(int64(time.Now().Second()))
	if rand.Intn(20) == 1 {
		reply := app.client.LPush("Errors", message)
		if err := reply.Err(); err != nil {
			fmt.Println("ERROR occured while LOGGING the incorrect message")
		} else {
			fmt.Println("The following message is INCORRECT: ", message)
		}
	} else {
		fmt.Println("The following message has been SUCCESSFULLY verified:", message)
	}
}
