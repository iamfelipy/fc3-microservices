package gateway

import "github.com/iamfelipy/fc3-microservices/wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
