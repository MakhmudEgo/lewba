package queue

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

const (
	SEND_ACK_EMAIL = "send-ack-email"
)

type Rabbit struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbit() *Rabbit {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	// todo::handler::register::need change
	if err != nil {
		log.Fatalln(err)
	}

	ch, err := conn.Channel()
	// todo::handler::register::need change
	if err != nil {
		log.Fatalln(err)
	}
	return &Rabbit{Conn: conn, Ch: ch}
}
