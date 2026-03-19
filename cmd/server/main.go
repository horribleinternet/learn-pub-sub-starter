package main

import (
	"fmt"
	"learn-pub-sub-starter/internal/pubsub"
	"learn-pub-sub-starter/internal/routing"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	connectstr := "amqp://guest:guest@localhost:5672/"
	connect, err := amqp.Dial(connectstr)
	if err != nil {
		fmt.Println("Peril server starting failed: ", err.Error())
		return
	}
	defer connect.Close()
	fmt.Println("Peril server started.")
	ch, err := connect.Channel()
	if err != nil {
		fmt.Println("Peril server unable to open channel: ", err.Error())
		return
	}
	err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
	if err != nil {
		fmt.Println("Error sending test message: ", err.Error())
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Interrupt received; Peril server shutting down.")

}
