
projeto: digital wallet full cycle

#### arquitetura: 
	- c4-model
	- para visualizar o restante usar o preview plantuml
![C4-Model Contexto](./docs/images/c4-model-context.png)  
description: context

#### components:  
	- microservice: wallet-core  
		- focado em clientes, contas e transações
		- stack: go  
		- como instalar tools go?  
			- ctrl + shift + p  
			- go: install/update tools  
			- instalar tudo  
	
	- package: utils-package
		- contem pacotes que podem ser compartilhados
		- stack: go  
		- como instalar tools go?  
			- ctrl + shift + p  
			- go: install/update tools  
			- instalar tudo  
		

#### ajuste para visualizar PlantUML no VSCode:
	- metodologia: C4 MODEL  
	- documentação em /docs

	1. Instale a extensão "PlantUML" (jebbs) no VSCode.

	2. Abra as configurações da extensão PlantUML:
		- Defina "Plantuml: Render" como `PlantUMLServer`
		- Defina "Plantuml: Server" como `http://plantuml-c4model:8080`

	3. Abra um arquivo `.puml` e utilize o preview do PlantUML no VSCode.

#### como usar o ambiente localmente? 
	- subir ambiente 
		- docker compose -f docker/docker-compose.yml up 
	- subir ambiente, apos atualizar docker file 
		- docker compose -f docker/docker-compose.yml up --build 
	- acessar microservice wallet-core 
		- docker-compose -f  docker/docker-compose.yml exec -it microservice-wallet-core bash 

#### como subir o ambiente usando devcontainer no vscode ou cursor? 
	- instalar extensão DevContainer
	- ctrl + shift + p  
	- escolher a opção "Dev Containers: Reopen in Container" 

