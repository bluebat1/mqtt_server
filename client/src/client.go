package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "127.0.0.1"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("client-1")
	opts.SetUsername("c1")
	opts.SetPassword("c1pd")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	client.Subscribe("topic_1", 1, func(c mqtt.Client, m mqtt.Message) {
		log.Println(m)
	})

	t := client.Publish("topic_1", 1, false, "hello")
	// <-t.Done()
	_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
	if t.Error() != nil {
		log.Panicln(t.Error()) // Use your preferred logging technique (or just fmt.Printf)
	}

	log.Println("end...")

	time.Sleep(time.Second * 2)
}
