package main

import (
	"github.com/nats-io/go-nats"
	"time"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
)

type dedicRequest struct {
	Hash   string
	Entity string
	Action string
	Id     int
}

func main (){

	nc, _ := nats.Connect("nats://0.0.0.0:4222") // change on real when docker enabled
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)


	c.QueueSubscribe("Dedek.*", "dedic_workers", func(req *dedicRequest) {
		logrus.Printf("Received a Dedek request: %s, to %s, reply to %s ", req.Action, req.Entity, req.Hash)

		rand.Seed(time.Now().Unix())
		myrand := random(1, 9999)

		time.Sleep(7 * time.Second)

		logrus.Printf("Replying to request with id: %d, hash: %s ", myrand, req.Hash)
		message := "All done. Sleep calm. ID: "+ strconv.Itoa(myrand)+" "
		logrus.Printf("message : %s \n", message)

		c.Publish(req.Hash, message)
	})


	c.Subscribe("foo", func(s string) {
		logrus.Printf("Received a message 'foo' service 2: %s", s)
	})

	c.QueueSubscribe("Some.more.>", "job_workers", func(m *nats.Msg) {
		logrus.Printf("Received some: %s\n", m.Data)
	})

	for{
		repeat_v(3*time.Second)
	}
}

func repeat_v(d time.Duration) {
	for range time.Tick(d) {
		logrus.Print("waiting......")
	}
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}
