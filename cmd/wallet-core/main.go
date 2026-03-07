package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamfelipy/fc3-microservices/internal/database"
	"github.com/iamfelipy/fc3-microservices/pkg/uow"

	"github.com/iamfelipy/fc3-microservices/internal/usecase/create_account"
	"github.com/iamfelipy/fc3-microservices/internal/usecase/create_client"
	"github.com/iamfelipy/fc3-microservices/internal/usecase/create_transaction"

	"github.com/iamfelipy/fc3-microservices/internal/web"
	"github.com/iamfelipy/fc3-microservices/internal/web/webserver"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iamfelipy/fc3-microservices/internal/event"
	"github.com/iamfelipy/fc3-microservices/internal/event/handler"
	"github.com/iamfelipy/fc3-microservices/pkg/events"
	"github.com/iamfelipy/fc3-microservices/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	// transactionDb := database.NewTransactionDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	configMap := ckafka.ConfigMap{
		//  lista inicial de brokers para conectar o cliente ao cluster.
		"bootstrap.servers": "kafka:29092",
		// so deve ser passada para consumidores
		// identifica o grupo de consumidores Kafka, útil para gerenciamento de consumidores e balanceamento de mensagens.
		"group.id": "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	transactionCreatedEvent := event.NewTransactionCreated()
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server started on port 8080")
	webserver.Start()
}
