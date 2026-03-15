import { FastifyInstance } from 'fastify'
import { getAccountBalance } from './account-balance'

export async function balancesRoutes(app: FastifyInstance) {
  app.get('/balances/:accountId', getAccountBalance)
}
