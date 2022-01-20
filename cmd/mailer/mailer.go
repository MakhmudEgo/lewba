package main

import (
	"crypto/tls"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("No .env file found")
	}
}

func main() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	//todo::
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"send-ack-email", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalln(err)
	}
	msgs, err := ch.Consume(
		q.Name,   // queue
		"mailer", // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	//err = ch.Qos(1, 1, false)
	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)
	//todo::mailer::add context and bot
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s | type: %s", d.Body, d.ContentType)
			err = send(string(d.Body))
			if err == nil {
				d.Ack(false)
			} else {
				d.Nack(false, true)
				println("nack sleep")
				time.Sleep(10 * time.Second)
			}
			println("~~~~√~~~~~~")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func send(addr string) error {
	m := gomail.NewMessage()
	mail := os.Getenv("EMAIL_CONFIRM_MAIL")
	m.SetHeader("From", mail)
	m.SetAddressHeader("From", mail, "Matcha")
	m.SetHeader("To", addr)

	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	//todo::create table for confirm emails
	uu, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	m.SetHeader("Subject", "Confirm email to matcha! (for test)")

	m.SetBody("text/plain", uu.String())
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), smtpPort, mail, os.Getenv("EMAIL_CONFIRM_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
