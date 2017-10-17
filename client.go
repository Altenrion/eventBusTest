package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"time"
)

func main (){

	nc, _ := nats.Connect("nats://0.0.0.0:4222") // change on real when docker enabled


	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message 'foo' service 1: %s\n", string(m.Data))
	})

	nc.QueueSubscribe("Some.more.>", "job_workers", func(m *nats.Msg) {
		fmt.Printf("Received some: %s\n", string(m.Data))
	})


	for{
		repeat_d(3*time.Second)
	}
}

func repeat_d(d time.Duration) {
	for range time.Tick(d) {
		fmt.Print("yeap, waiting... \n")
	}
}
