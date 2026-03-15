import { prisma } from '@/modules/balance/infra/repository/prisma/prisma.service'
import { PrismaAccountRepository } from '../../repository/prisma/repositories/account.prisma.repository'
import { UpdateBalanceUseCase } from '@/modules/balance/domain/usecase/update-balance.usecase'

export function makeUpdateBalanceUseCase(): UpdateBalanceUseCase {
  const accountRepository = new PrismaAccountRepository(prisma)
  const updateBalanceUseCase = new UpdateBalanceUseCase(accountRepository)
  return updateBalanceUseCase
}
