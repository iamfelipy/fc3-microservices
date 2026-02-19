
projeto: digital wallet full cycle

#### microservices:
   - wallet-core

#### ajuste para visualizar PlantUML no VSCode:
   - metodologia: C4 MODEL  

   1. Instale a extensão "PlantUML" (jebbs) no VSCode.

   2. Abra as configurações da extensão PlantUML:
      - Defina "Plantuml: Render" como `plantumlserver`
      - Defina "Plantuml: Server" como `http://localhost:8080`

   3. Execute `docker-compose up` para iniciar o servidor PlantUML.

   4. Abra um arquivo `.puml` e utilize o preview do PlantUML no VSCode.


#### como usar o ambiente localmente? 
   - subir ambiente 
      - docker compose -f docker/docker-compose.yml up 
   - subir ambiente, apos atualizar docker file 
      - docker compose -f docker/docker-compose.yml up --build 
   - acessar microservice wallet-core 
      - docker-compose -f  docker/docker-compose.yml exec -it microservice-wallet-core bash 

#### como subir o ambiente usando devcontaine(developer cloud) no vscode ou cursor? 
   - ctrl + shift + p  
   - escolher a opção "Dev Containers: Rebuild Without Cache and Reopen in Container" 

