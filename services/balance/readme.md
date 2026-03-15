# microservice: balance
 
- Retorna apenas o saldo atual do cliente  
- stack: typescript - nodejs  

#### setup Rápido
```bash

# docker-compose já inicializa tudo, talvez não precise usar

# instalar dependencias
npm i

# Configure variáveis de ambiente
cp .env.example .env

# Edite .env com suas credenciais

# Inicie a aplicação
npm run start:dev
```

#### comandos uteis

```bash
# Os principais comandos do Prisma (v5) para trabalhar com múltiplos contexts
  # Exemplos considerando o schema.prisma em:
  # ./src/modules/balance/infra/repository/prisma/schema.prisma

  # Inicializa configuração do Prisma para o caminho customizado
    npx prisma init --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # Gera o cliente Prisma para este schema
    npx prisma generate --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # Gera uma nova migration e aplicar no banco (substitua NOME_DA_MIGRATION pelo nome desejado)
    # banco preciso estar ativo
    npx prisma migrate dev --name NOME_DA_MIGRATION --schema ./src/modules/balance/infra/repository/prisma/schema.prisma 
    npx prisma migrate dev --name create_accounts  --schema ./src/modules/balance/infra/repository/prisma/schema.prisma
  # Gera uma nova migration sem executar (substitua NOME_DA_MIGRATION pelo nome desejado)
    # banco preciso estar ativo
    npx prisma migrate dev --create-only --name NOME_DA_MIGRATION --schema ./src/modules/balance/infra/repository/prisma/schema.prisma 
    npx prisma migrate dev --create-only --name create_accounts  --schema ./src/modules/balance/infra/repository/prisma/schema.prisma
  # Executa migrações para desenvolvimento nesse contexto
    npx prisma migrate dev --schema ./src/modules/balance/infra/repository/prisma/schema.prisma 
  # Aplica migrações em produção
    npx prisma migrate deploy --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # Atualiza o banco de dados conforme o schema sem criar migrações
    npx prisma db push --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # Abre a interface visual para este contexto do banco de dados
    npx prisma studio --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # Formata o schema.prisma deste contexto
    npx prisma format --schema ./src/modules/balance/infra/repository/prisma/schema.prisma  
  # aplicar seed
    npx tsx src/modules/balance/infra/repository/prisma/seed.ts
  # Dica: Sempre use o parâmetro --schema para garantir que está trabalhando no schema do módulo correto.

# postgresql
  docker exec -it workspace-fc3-microservices-postgresql-1 bash
  psql -U root -d balance
  root
  SELECT * FROM accounts;
  UPDATE accounts SET id = '3f06b755-b302-4a61-a8de-55b9a4dc14a1', client_id = '44277e46-431a-4258-b559-55be47be593b', balance = 750, created_at = '2026-03-05' WHERE id = '1';
  UPDATE accounts SET id = 'cc8f516e-4d6c-4bd6-a22e-28010b74047e', client_id = 'a7ca7003-f312-4cd6-8919-282b6a1f9543', balance = 250, created_at = '2026-03-05' WHERE id = '2';

# typescript
  npx tsc --noEmit src/main.ts
```

