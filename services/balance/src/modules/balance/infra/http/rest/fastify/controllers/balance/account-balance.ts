import { z } from 'zod'
import { FastifyReply, FastifyRequest } from 'fastify'
import { GetBalanceUseCase } from '@/modules/balance/domain/usecase/get-balance.usecase'
import { PrismaAccountRepository } from '@/modules/balance/infra/repository/prisma/repositories/account.prisma.repository'
import { prisma } from '@/modules/balance/infra/repository/prisma/prisma.service'

export async function getAccountBalance(
  request: FastifyRequest,
  reply: FastifyReply,
) {
  const accountBalanceParamsSchema = z.object({
    accountId: z.string(),
  })

  const { accountId } = accountBalanceParamsSchema.parse(request.params)

  const accountRepository = new PrismaAccountRepository(prisma)
  const getBalanceUseCase = new GetBalanceUseCase(accountRepository)
  const balance = await getBalanceUseCase.execute({
    accountId,
  })

  return reply.status(200).send(balance)
}
