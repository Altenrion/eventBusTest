package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"time"
)

func main (){

	nc, _ := nats.Connect(nats.DefaultURL) // change on real when docker enabled
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	for{
		repeat_d(2*time.Second)
	}
}

func repeat_d(d time.Duration) {
	for range time.Tick(d) {

		fmt.Print("yeap, waiting... \n")

		//con.Publish("foo", []byte("Hello World"))
	}
}
