import { ZodError } from 'zod'
import { env } from '../../env'
import { makeUpdateBalanceUseCase } from '../../factory/usecase/update-balance.usecase.factory'
import { KafkaConsumerAdapter } from './kafka-consumer.adapter'
import { BalanceUpdatedHandler } from '@/modules/balance/domain/event/handlers/balance-updated-sync-db.handler'

const consumer = new KafkaConsumerAdapter({
  clientId: 'balance-service',
  brokers: [...env.KAFKA_BROKERS],
  // Com groupId, as mensagens são divididas entre os consumers do grupo.
  groupId: 'balance-group',
  topics: ['balances'],
})

consumer
  .init(async (message) => {
    if (!message.value) return
    const value = JSON.parse(message.value)

    switch (message.topic) {
      case 'balances': {
        const updateBalanceUseCase = makeUpdateBalanceUseCase()
        const balanceUpdatedHandler = new BalanceUpdatedHandler(
          updateBalanceUseCase,
        )
        await balanceUpdatedHandler.handle(value)

        break
      }
      default:
        // Topic desconhecido, ignorar ou logar
        break
    }
  })
  .then(() => {
    console.log('Kafka consumer rodando')
  })
  .catch((error) => {
    if (error instanceof ZodError) {
      console.log(error.format())
      // Poderia adicionar lógica extra para ZodError aqui, se necessário
    } else {
      // Apenas log de erro, seguindo app.ts
      console.error(error)
      // Aqui futuramente pode-se adicionar integração com ferramentas externas de logging
    }
  })
