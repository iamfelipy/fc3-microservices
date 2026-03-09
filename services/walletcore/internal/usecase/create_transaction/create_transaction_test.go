package create_transaction

import (
	"context"
	"testing"

	"github.com/iamfelipy/fc3-microservices/services/walletcore/internal/entity"
	"github.com/iamfelipy/fc3-microservices/services/walletcore/internal/event"
	"github.com/iamfelipy/fc3-microservices/services/walletcore/internal/usecase/mocks"
	"github.com/iamfelipy/fc3-microservices/services/walletcore/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// substituido por mockOuw
// type TransactionGatewayMock struct {
// 	mock.Mock
// }

// func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
// 	args := m.Called(transaction)
// 	return args.Error(0)
// }

// type AccountGatewayMock struct {
// 	mock.Mock
// }

// func (m *AccountGatewayMock) Save(account *entity.Account) error {
// 	args := m.Called(account)
// 	// se o on return que monta args, não tiver algo do tipo Error, retorna nil
// 	return args.Error(0)
// }

// func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entity.Account), args.Error(1)
// }

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "j@j.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "j@j2.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	// substituido por mockOuw
	// mockAccount := &AccountGatewayMock{}
	// mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	// mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	// mockTransaction := &TransactionGatewayMock{}
	// mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	transactionEvent := event.NewTransactionCreated()
	balanceEvent := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(
		mockUow,
		dispatcher,
		transactionEvent,
		balanceEvent,
	)

	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
	// substituido por mockOuw
	//assert.NotEmpty(t, output.ID)
	// mockAccount.AssertExpectations(t)
	// mockTransaction.AssertExpectations(t)
	// mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	// mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
