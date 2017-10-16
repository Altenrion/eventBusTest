package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"time"
)

func main (){

	nc, _ := nats.Connect(nats.DefaultURL) // change on real when docker enabled
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received new message: %s\n", string(m.Data))
	})

	for{
		repeat_v(15*time.Second, nc)
	}
}

func repeat_v(d time.Duration, con *nats.Conn) {
	for range time.Tick(d) {

		fmt.Print("yeap, waiting... \n")

		con.Publish("foo.back", []byte("Where is my new data, m?"))
	}
}
