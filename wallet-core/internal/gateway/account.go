package gateway

import "github.com/iamfelipy/fc3-microservices/wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
