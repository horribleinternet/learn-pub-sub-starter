package main

import (
	"fmt"
	"learn-pub-sub-starter/internal/gamelogic"
	"learn-pub-sub-starter/internal/pubsub"
	"learn-pub-sub-starter/internal/routing"

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
	ch, err := connect.Channel()
	if err != nil {
		fmt.Println("Peril server unable to open channel: ", err.Error())
		return
	}
	gamelogic.PrintServerHelp()
	looping := true
	for looping {
		inputs := gamelogic.GetInput()
		if len(inputs) < 1 {
			continue
		}
		switch inputs[0] {
		case "pause":
			fmt.Println("Pausing...")
			err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
			if err != nil {
				fmt.Println("Error sending pause message: ", err.Error())
			}
		case "resume":
			fmt.Println("Resuming...")
			err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: false})
			if err != nil {
				fmt.Println("Error sending resume message: ", err.Error())
			}
		case "quit":
			fmt.Println("Exiting...")
			looping = false
		default:
			fmt.Println("Unrecognized command.")
		}
	}
	//signalChan := make(chan os.Signal, 1)
	//signal.Notify(signalChan, os.Interrupt)
	//<-signalChan
	//fmt.Println("Interrupt received; Peril server shutting down.")

}
