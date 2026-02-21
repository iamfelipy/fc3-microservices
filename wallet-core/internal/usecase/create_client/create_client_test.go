package createclient

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

// fingir uma conexão com banco, que tem o metodo save
// fingir que estou indo no banco de dados

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "John Doe",
		Email: "jej",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "jej", output.Email)
	// garante que o save foi chamado
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
