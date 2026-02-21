package createaccount

import (
	"testing"

	"github.com/iamfelipy/fc3-microservices/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	// traz varios metodos proprios pra dentro
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	// verifica se o mock foi chamado com o client
	args := m.Called(client)
	// retorna o que esta na posição 0 do return e converte para erro
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	// Sim. O .() é conversão de tipo (type assertion), não acesso a campo de struct.
	// .Nome → acessa campo de struct
	// .(*entity.Client) → converte interface{} pro tipo real
	return args.Get(0).(*entity.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	// se o on return que monta args, não tiver algo do tipo Error, retorna nil
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jej")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
