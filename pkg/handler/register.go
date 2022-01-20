package handler

import (
	"io/ioutil"
	"log"
	"matcha/pkg/queue"
	"matcha/pkg/repository"
	"matcha/pkg/service"
	"net/http"
)

type register struct {
	db     repository.UserDB
	rabbit *queue.Rabbit
}

func Register(db repository.UserDB) *register {
	return &register{db: db, rabbit: queue.NewRabbit()}
}

func (reg *register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	// todo::handler::register::придумать алгоритм обработки ошибок user/server
	if err != nil {
		log.Fatal(err)
	}
	if _, err = service.Register(reg.db, reg.rabbit).Exec(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("this is user exist"))
		log.Println(err)
	} else {
		w.WriteHeader(201)
	}
}
