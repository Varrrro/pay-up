package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"github.com/varrrro/pay-up/internal/consumer"
	"github.com/varrrro/pay-up/internal/publisher"
	"github.com/varrrro/pay-up/internal/tmicro"
)

func main() {
	// Create channel to listen for OS signals
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, os.Interrupt, os.Kill)

	// Create context that can be cancelled
	ctx, cfunc := context.WithCancel(context.Background())
	defer cfunc()

	// Open AMQP connection
	conn, err := amqp.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Open database connection
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()

	// Create data manager and handlers
	tm := tmicro.NewManager(db)

	// Create AMQP publisher
	p, err := publisher.New(conn, "transactions", "balance")
	if err != nil {
		log.Fatal(err)
	}

	// Create message handler
	h := tmicro.NewMessageHandler(tm, p)

	// Create AMQP consumer
	c, err := consumer.New(conn, "transactions", "management", "tmicro")
	if err != nil {
		log.Fatal(err)
	}

	c.Start(ctx, h) // start consumer

	<-sch // blocking until we receive a signal
}
