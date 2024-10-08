package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	ledTopic = "led"
	clientId = "go_mqtt_client"
)

func main() {
	broker := os.Getenv("MQTT_BROKER")
	if broker == "" {
		log.Fatal("MQTT_BROKER environment variable is not set")
	}

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting: %s", token.Error())
	}
	defer client.Disconnect(250)

	go func() {
		for {
			fmt.Print("Press Enter to send a message to the led topic...")
			fmt.Scanln()
			if token := client.Publish(ledTopic, 0, false, "LED ON"); token.Wait() && token.Error() != nil {
				log.Printf("Error publishing to %s: %s", ledTopic, token.Error())
			} else {
				log.Printf("Message sent to topic: %s", ledTopic)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down MQTT client...")
}
