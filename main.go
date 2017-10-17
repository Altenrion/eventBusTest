package main

import (
	"fmt"
	"time"
	"github.com/nats-io/go-nats"
	"math/rand"
)

func main (){

	fmt.Print("Hello, World!\n")
	nc, _ := nats.Connect("nats://0.0.0.0:4222")

	nc.Subscribe("Error", func(m *nats.Msg) {
		fmt.Printf("Error: %s\n", string(m.Data))
		fmt.Print("God, I have to do something....\n")

	})

	repeat(5*time.Second, nc)
}


func repeat(d time.Duration, con *nats.Conn ) {
	for range time.Tick(d) {

		fmt.Print("------------------------------\n")
		fmt.Print("passing some data into bus... \n")

		subj, text := eventsProducer()
		con.Publish(string(subj), []byte(text))
	}
}

func eventsProducer() (string, string){
	events := make([][]string, 0)
	events = append(events,
		[]string{"Locked"," Connection to remote server was locked gracefully"},
		[]string{"Error","Some error occurred in our environment"},
		[]string{"New_User","New user was added in User_Service. React on that fact properly"},
		[]string{"Fatal","Oh my God! Fatal disaster in da house"},
		[]string{"KPI:new","New KPI indicator created"},
		[]string{"Bad","Oh yeah...Bad day"},
		[]string{"Some.more.event","Subscription on events"},
		[]string{"Some.more.sudden","Subscription on sudden"},
		[]string{"Some","First level subscriptions"},
		[]string{"foo","Just foo test"},
	)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	event := r.Intn(len(events))

	fmt.Printf("%s : %s  \n",  events[event][0], events[event][1])
	return events[event][0], events[event][1]
}