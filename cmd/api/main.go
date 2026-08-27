package main

import (
	"log"
	"net/http"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/messaging"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/config"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/controllers"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/database"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/repository"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/routes"
	"github.com/Ludimila-Araujo/labsupply-payment/internal/service"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	paymentRepository := repository.NewPostgresPaymentRepository(db)

	paymentService := service.NewPaymentService(
		paymentRepository,
	)

	rabbitMQ, err := messaging.NewRabbitMQ(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal(err)
	}

	defer rabbitMQ.Close()

	if err := rabbitMQ.DeclarePaymentQueue(); err != nil {
		log.Fatal(err)
	}

	messages, err := rabbitMQ.ConsumePayments()
	if err != nil {
		log.Fatal(err)
	}

	paymentConsumer := messaging.NewPaymentConsumer(paymentService)

	go paymentConsumer.Start(messages)

	paymentController := controllers.NewPaymentController(
		paymentService,
	)

	router := routes.Setup(
		paymentController,
	)

	log.Println("payment service running on :8081")

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
