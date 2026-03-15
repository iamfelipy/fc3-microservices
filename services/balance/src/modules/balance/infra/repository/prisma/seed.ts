import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

async function main() {
  // Criar contas (Account) - pois User está comentado no schema.prisma (refletindo schema atual)
  // IDs fixos para as contas, igual ao walletcore/seed/seed.go
  const client1ID = 'a7ca7003-f312-4cd6-8919-282b6a1f9543'
  const client2ID = '44277e46-431a-4258-b559-55be47be593b'
  const account1ID = '3f06b755-b302-4a61-a8de-55b9a4dc14a1'
  const account2ID = 'cc8f516e-4d6c-4bd6-a22e-28010b74047e'

  await prisma.account.upsert({
    where: { id: account1ID },
    update: {},
    create: {
      id: account1ID,
      clientId: client1ID,
      balance: 1000,
      // podemos colocar createdAt manual se schema pedir, mas Prisma costuma preencher automático
    },
  })

  await prisma.account.upsert({
    where: { id: account2ID },
    update: {},
    create: {
      id: account2ID,
      clientId: client2ID,
      balance: 500,
    },
  })

  console.log('Seed completed!')
}

main()
  .catch((e) => {
    console.error(e)
    process.exit(1)
  })
  .finally(async () => {
    await prisma.$disconnect()
  })
