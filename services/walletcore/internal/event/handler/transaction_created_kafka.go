package handler

import (
	"fmt"
	"sync"

	"github.com/iamfelipy/fc3-microservices/walletcore/pkg/events"
	"github.com/iamfelipy/fc3-microservices/walletcore/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

// wait group
// impedir que o processo principal encerre, enquanto as goroutines ainda estao trabalhando
func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	// avisa que a goroutine terminou
	defer wg.Done()
	// message, key, topic
	h.Kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
}
