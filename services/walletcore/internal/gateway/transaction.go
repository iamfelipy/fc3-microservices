package gateway

import "github.com/iamfelipy/fc3-microservices/services/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
