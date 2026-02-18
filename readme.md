metodologia: C4 MODEL  
projeto: digital wallet full cycle

ajuste para visualizar PlantUML no VSCode:

1. Instale a extensão "PlantUML" (jebbs) no VSCode.

2. Abra as configurações da extensão PlantUML:
   - Defina "Plantuml: Render" como `plantumlserver`
   - Defina "Plantuml: Server" como `http://localhost:8080`

3. Execute `docker-compose up` para iniciar o servidor PlantUML.

4. Abra um arquivo `.puml` e utilize o preview do PlantUML no VSCode.