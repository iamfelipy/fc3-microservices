package gateway

import "github.com/iamfelipy/fc3-microservices/services/walletcore/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
