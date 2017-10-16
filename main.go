package main

import (
	"fmt"
	"time"
	"github.com/nats-io/go-nats"
)

func main (){

	fmt.Printf("Hello, World!\n")
	nc, _ := nats.Connect("nats://0.0.0.0:4222")

	nc.Subscribe("foo.back", func(m *nats.Msg) {
		fmt.Printf("Reply: %s\n", string(m.Data))
	})

	repeat(5*time.Second, nc)
}


func repeat(d time.Duration, con *nats.Conn ) {
	for range time.Tick(d) {

		fmt.Print("yeap, some time passed \n")

		con.Publish("foo", []byte("Hello World"))
	}
}
