package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"time"
)

func main (){

	nc, _ := nats.Connect("nats://0.0.0.0:4222") // change on real when docker enabled


	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Subscribe("Locked", func(m *nats.Msg) {
		fmt.Printf("Locked: %s\n", string(m.Data))
		fmt.Print("Publishing error event ...")

		nc.Publish("Error", []byte("All is locked"))

	})

	nc.Subscribe("Fatal", func(m *nats.Msg) {
		fmt.Printf("We are in danger: %s\n", string(m.Data))

		nc.Publish("Error", []byte("Run Fullish , Run!"))

	})

	nc.Subscribe("Some.*", func(m *nats.Msg) {
		fmt.Printf("New Some..: %s\n", string(m.Data))
		fmt.Printf("Well,... new Some event came : %s ", m.Subject)
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
