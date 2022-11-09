package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gustavohenriquess/golang-intensive/internal/order/infra/database"
	"github.com/gustavohenriquess/golang-intensive/internal/order/usecase"
	"github.com/gustavohenriquess/golang-intensive/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery) //Channel to receive messages
	go rabbitmq.Consume(ch, out)    //T2

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)

		fmt.Println(outputDTO) //T1
		time.Sleep(400 * time.Millisecond)
	}
}
