package main

import (
	"fmt"
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
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Interrupt received; Peril server shutting down.")

}
