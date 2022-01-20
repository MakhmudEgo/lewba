package service

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"matcha/pkg/model"
	"matcha/pkg/queue"
	"matcha/pkg/repository"
	"runtime"
)

type IRegister interface {
}

type register struct {
	db     repository.UserDB
	rabbit *queue.Rabbit
}

func Register(db repository.UserDB, rabbit *queue.Rabbit) Service {
	return &register{db: db, rabbit: rabbit}
}

func (r *register) Do(user *model.User) (interface{}, error) {

	panic("implement me")
}

func (r *register) Exec(data []byte) (interface{}, error) {
	user := &model.User{}
	err := json.Unmarshal(data, &user)
	// todo::service::register::need change err
	if err != nil {
		return nil, err
	}
	// todo::service::register parsing and valid user

	// todo::service::register if user ok
	//todo::service::register::context
	go func() {
		if err = r.db.CreateUser(context.Background(), user); err == nil {
			go addQueue(r.rabbit, user.Email)
		}
	}()
	//todo:://service::register::имитация, как-будто все хорошо
	return nil, nil
}

func addQueue(rabbit *queue.Rabbit, email string) {
	q, err := rabbit.Ch.QueueDeclare(
		queue.SEND_ACK_EMAIL, // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Println(err, "Failed to declare a queue")
		log.Println(runtime.Caller(0))
	}

	err = rabbit.Ch.Publish(
		"mailer", // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(email),
		})
	if err != nil {
		log.Println(err, "Failed to publish a queue")
		log.Println(runtime.Caller(0))
	}
}
