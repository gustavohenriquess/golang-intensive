package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
	// forever := make(chan bool)      //Channel to keep the program running
	go rabbitmq.Consume(ch, out) //T2

	qtdWorkers := 200

	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}
	// <-forever // Channel is never closed, so the program will never end

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecase.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(total)

	})
	http.ListenAndServe(":8080", nil)
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	for msg := range deliveryMessage {
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

		fmt.Printf("Worker %d has processed order %s\n", workerID, outputDTO.ID) //T1
		time.Sleep(1 * time.Second)
	}
}
