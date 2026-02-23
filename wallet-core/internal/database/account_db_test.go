package database

import (
	"database/sql"
	"testing"

	"github.com/iamfelipy/fc3-microservices/wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	_, err = s.db.Exec(`CREATE TABLE clients (
		id VARCHAR(255), 
		name VARCHAR(255), 
		email VARCHAR(255), 
		created_at DATE,
		updated_at DATE
	)`)
	s.Nil(err)
	_, err = s.db.Exec(`CREATE TABLE accounts (
		id VARCHAR(255), 
		client_id VARCHAR(255), 
		balance INTEGER, 
		created_at DATE, 
		updated_at DATE
	)`)
	s.Nil(err)
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) createClient() *entity.Client {
	client, err := entity.NewClient("Test", "test@example.com")
	s.Nil(err)
	_, err = s.db.Exec(`INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)`,
		client.ID, client.Name, client.Email, client.CreatedAt)
	s.Nil(err)
	return client
}

func (s *AccountDBTestSuite) TestSave() {
	// não configurei verificação de integridade referencial
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec(`INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)`,
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
}
