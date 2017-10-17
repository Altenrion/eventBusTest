package main

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"time"
)

func main (){

	nc, _ := nats.Connect("nats://0.0.0.0:4222") // change on real when docker enabled


	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message 'foo' service 2: %s\n", string(m.Data))
	})

	nc.QueueSubscribe("Some.more.>", "job_workers", func(m *nats.Msg) {
		fmt.Printf("Received some: %s\n", string(m.Data))
	})

	//nc.Subscribe("Some.*.*", func(m *nats.Msg) {
	//	fmt.Printf("New Some..: %s\n", string(m.Data))
	//	fmt.Print("Well,... but who what is this... some about ??? \n")
	//
	//	nc.Publish(m.Reply, []byte("new Some event came:"+m.Subject))
	//
	//})


	for{
		repeat_v(3*time.Second, nc)
	}
}

func repeat_v(d time.Duration, con *nats.Conn) {
	for range time.Tick(d) {

		fmt.Print("yeap, waiting... \n")

		con.Publish("foo.back", []byte("Where is my new data, m? \n"))
	}
}
