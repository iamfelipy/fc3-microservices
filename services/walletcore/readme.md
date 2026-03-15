# microservice: wallet-core  
 
- focado em clientes, contas e transações
- stack: go  
- package: pkg
	- contem pacotes que podem ser compartilhados


#### como instalar tools go no vscode?  
- instalar a extensão go
	- dev container já traz instalada
- ctrl + shift + p  
- go: install/update tools  
	- instalar tudo  

#### executar testes go
```bash
// apartir da raiz
go test ./...
// apartir de um pacote
go test ../...
```

#### observações
- extension vscode: REST Client
	- usada para testar a api
	- walletcore/api/client.http


#### comandos uteis para diferentes components
```bash
# docker-compose já inicializa tudo, talvez não precise usar

# iniciar webservice
	go run cmd/wallet-core/main.go

# mysql
	docker ps
	docker exec -it workspace-fc3-microservices-mysql-1  bash
	mysql -uroot -p wallet
	root

	CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at date);
	CREATE TABLE accounts (id VARCHAR(255) PRIMARY KEY, client_id VARCHAR(255), balance INT, created_at date);
	CREATE TABLE transactions (id VARCHAR(255) PRIMARY KEY, account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount INT, created_at date);

	select * from clients;
	select * from accounts;
	DROP TABLE clients;
	DROP TABLE accounts;
	DROP TABLE transactions;
	update accounts set balance=100 where id="3f06b755-b302-4a61-a8de-55b9a4dc14a1";

```