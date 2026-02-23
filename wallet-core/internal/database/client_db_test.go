package database

import (
	"database/sql"
	"testing"

	// injeta o sqlite3 no database/sql
	"github.com/iamfelipy/fc3-microservices/wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	//  Ao embutir suite.Suite, todos os métodos e campos públicos dela ficam acessíveis direto pelo ClientDBTestSuite, funcionando como se fossem "parte" dele.
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	s.db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

// inicializar o ambiente da suíte de testes
func TestClientDBTestSuite(t *testing.T) {
	// new: criar uma instancia da struct sem precisa inicializar variavel
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "Test",
		Email: "j@j.com",
	}

	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John", "j@j.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
