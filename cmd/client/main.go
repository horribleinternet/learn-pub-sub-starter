package main

import (
	"fmt"
	"learn-pub-sub-starter/internal/gamelogic"
	"learn-pub-sub-starter/internal/pubsub"
	"learn-pub-sub-starter/internal/routing"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	connectstr := "amqp://guest:guest@localhost:5672/"
	connect, err := amqp.Dial(connectstr)
	if err != nil {
		fmt.Println("Peril server client failed: ", err.Error())
		return
	}
	defer connect.Close()
	fmt.Println("Peril client started.")
	username, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Println("Peril client error reading user name: ", username, err.Error())
		return
	}
	_, _, err = pubsub.DeclareAndBind(connect, routing.ExchangePerilDirect, routing.PauseKey+"."+username, routing.PauseKey, pubsub.Transient)
	if err != nil {
		fmt.Println("Error binding queue: ", err.Error())
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Interrupt received; Peril server shutting down.")
}
