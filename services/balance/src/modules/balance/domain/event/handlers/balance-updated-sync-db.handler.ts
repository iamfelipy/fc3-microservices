import EventHandlerInterface from '@/modules/@shared/event/event-handler.interface'
import { UpdateBalanceUseCase } from '@/modules/balance/domain/usecase/update-balance.usecase'
import BalanceUpdatedEvent from '../balance-updated.event'
import { z, ZodError } from 'zod'

export class BalanceUpdatedHandler
  implements EventHandlerInterface<BalanceUpdatedEvent>
{
  constructor(private readonly useCase: UpdateBalanceUseCase) {}

  async handle(data: unknown): Promise<void> {
    const payload = this.validateAndExtractPayload(data)

    await this.useCase.execute({
      accountId: payload.account_id_from,
      balance: payload.balance_account_id_from,
    })
    await this.useCase.execute({
      accountId: payload.account_id_to,
      balance: payload.balance_account_id_to,
    })
  }

  validateAndExtractPayload(data: unknown) {
    const BalanceUpdatedEventDTOSchema = z.object({
      Name: z.literal('BalanceUpdated'),
      Payload: z.object({
        account_id_from: z.string(),
        account_id_to: z.string(),
        balance_account_id_from: z.number(),
        balance_account_id_to: z.number(),
      }),
    })

    type BalanceUpdatedEventDTO = z.infer<typeof BalanceUpdatedEventDTOSchema>

    const validation = BalanceUpdatedEventDTOSchema.safeParse(data)
    if (!validation.success) {
      const zodError = new ZodError(validation.error.issues)
      console.error(
        "Payload inválido ao processar BalanceUpdatedHandler: veja 'issues' abaixo.",
        validation.error.format(),
      )
      throw zodError
    }

    const event = new BalanceUpdatedEvent<BalanceUpdatedEventDTO['Payload']>()
    event.setPayload(validation.data.Payload)

    return event.getPayload()
  }
}
