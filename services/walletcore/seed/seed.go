package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "fc3-microservices-mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Cria tabelas se não existirem e exclui registros antigos das tabelas para evitar conflitos ao rodar o seed múltiplas vezes

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			created_at date
		)
	`)
	if err != nil {
		panic(fmt.Sprintf("Error creating clients table: %v", err))
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id VARCHAR(255) PRIMARY KEY,
			client_id VARCHAR(255),
			balance INT,
			created_at date
		)
	`)
	if err != nil {
		panic(fmt.Sprintf("Error creating accounts table: %v", err))
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id VARCHAR(255) PRIMARY KEY,
			account_id_from VARCHAR(255),
			account_id_to VARCHAR(255),
			amount INT,
			created_at date
		)
	`)
	if err != nil {
		panic(fmt.Sprintf("Error creating transactions table: %v", err))
	}

	// Limpar as tabelas antes de inserir os dados seed
	_, err = db.Exec("DELETE FROM transactions")
	if err != nil {
		panic(fmt.Sprintf("Error deleting from transactions: %v", err))
	}
	_, err = db.Exec("DELETE FROM accounts")
	if err != nil {
		panic(fmt.Sprintf("Error deleting from accounts: %v", err))
	}
	_, err = db.Exec("DELETE FROM clients")
	if err != nil {
		panic(fmt.Sprintf("Error deleting from clients: %v", err))
	}

	client1ID := "a7ca7003-f312-4cd6-8919-282b6a1f9543"
	client2ID := "44277e46-431a-4258-b559-55be47be593b"
	account1ID := "3f06b755-b302-4a61-a8de-55b9a4dc14a1"
	account2ID := "cc8f516e-4d6c-4bd6-a22e-28010b74047e"

	// Inserir clients com IDs e dados fixos - se já existir, não dá erro
	_, err = db.Exec(`INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, NOW())
					  ON DUPLICATE KEY UPDATE name=VALUES(name), email=VALUES(email)`, client1ID, "John Doe", "j@j.com")
	if err != nil {
		panic(fmt.Sprintf("Error inserting client 1: %v", err))
	}
	_, err = db.Exec(`INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, NOW())
					  ON DUPLICATE KEY UPDATE name=VALUES(name), email=VALUES(email)`, client2ID, "Jane Doe", "jane@j.com")
	if err != nil {
		panic(fmt.Sprintf("Error inserting client 2: %v", err))
	}

	// Inserir accounts com IDs fixos - se já existir, não dá erro
	_, err = db.Exec(`INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, NOW())
					  ON DUPLICATE KEY UPDATE balance=VALUES(balance), client_id=VALUES(client_id)`, account1ID, client1ID, 1000)
	if err != nil {
		panic(fmt.Sprintf("Error inserting account for client 1: %v", err))
	}
	_, err = db.Exec(`INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, NOW())
					  ON DUPLICATE KEY UPDATE balance=VALUES(balance), client_id=VALUES(client_id)`, account2ID, client2ID, 500)
	if err != nil {
		panic(fmt.Sprintf("Error inserting account for client 2: %v", err))
	}

	fmt.Println("Seed completed!")
}
