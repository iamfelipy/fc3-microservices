import { env } from '@/modules/balance/infra/env'
import { PrismaClient } from '@prisma/client'

export class PrismaService
  extends PrismaClient
{
  constructor() {
    super({
      log: env.NODE_ENV === 'dev' ? ['query'] : [],
    })
  }
}


export const prisma = new PrismaService()

