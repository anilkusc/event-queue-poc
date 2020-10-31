package main

import (
	"fmt"
	"time"

	gen "github.com/anilkusc/golang-random-string-generator/generator"
	"github.com/streadway/amqp"
)

func main() {
	for {
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

		q, err := ch.QueueDeclare(
			"GenerateString", //name
			false,            // durable
			false,            // autoDelete
			false,            // exclusive
			false,            // no wait
			nil,
		)
		fmt.Println(q)
		if err != nil {
			fmt.Println(err)
		}

		err = ch.Publish(
			"",               //exchange
			"GenerateString", //key
			false,            //mandatory
			false,            //immediate
			amqp.Publishing{ // message
				ContentType: "text/plain",
				Body:        []byte(gen.Generate(10)),
			},
		)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Published Message to Queue")
		time.Sleep(2 * time.Minute)
	}
}
