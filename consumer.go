package main

import (
	"fmt"
	"net/http"

	"github.com/streadway/amqp"
)

var pass string = ""

func ListenChannel() {
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.2:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"GenerateString",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			pass = string(d.Body)

		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever

}

func GetString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", pass)
}

func main() {
	go ListenChannel()
	http.HandleFunc("/", GetString)
	http.ListenAndServe(":8080", nil)
}
