package main

import (
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"

	"crypto/md5"
	"encoding/hex"
)

type dedicRequest struct {
	Hash   string
	Entity string
	Action string
	Id     int
}

func main() {

	logrus.Print("Hello, World!\n")
	nc, _ := nats.Connect("nats://0.0.0.0:4222")
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	defer c.Close()

	repeat(5*time.Second, c)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func repeat(d time.Duration, con *nats.EncodedConn) {
	for range time.Tick(d) {

		logrus.Print("---------------------------------------- ")
		logrus.Print("passing data into shared message queue | ")

		timeStamp := time.Now()

		if timeStamp.Second() > 45 {

			queueHash := GetMD5Hash(timeStamp.String())
			logrus.Info("-------------------------------------------- ")
			logrus.Infof("Created hash : %s \n", queueHash)

			request := &dedicRequest{Hash: queueHash, Entity: "User", Action: "Delete", Id: 111}

			logrus.Infof("Publishing type  Dedek.typed ")
			con.Publish(string("Dedek.typed"), request)

			logrus.Infof("Subscribing for %s topic ", queueHash)
			con.Subscribe(queueHash, func(m *nats.Msg) {
				logrus.Infof("Received some: %s", string(m.Data))
				m.Sub.Unsubscribe()
				logrus.Infof("Must be unsubscribed from %s with %s \n", m.Sub.Subject, m.Data)
			})

		}

		subj, text := eventsProducer()
		con.Publish(string(subj), []byte(text))
	}
}

func eventsProducer() (string, string) {
	events := make([][]string, 0)
	events = append(events,
		[]string{"Locked", " Connection to remote server was locked gracefully"},
		[]string{"Some.more.event", "Subscription on events"},
		[]string{"Some.more.sudden", "Subscription on sudden"},
		[]string{"Some.s", "First level subscriptions"},
		[]string{"foo", "Just foo test"},
	)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	event := r.Intn(len(events))

	logrus.Printf("%s : %s ", events[event][0], events[event][1])
	return events[event][0], events[event][1]
}
