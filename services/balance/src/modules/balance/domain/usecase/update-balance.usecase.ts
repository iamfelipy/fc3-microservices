import { AccountRepository } from '@/modules/balance/domain/repository/account.repository'

export interface UpdateBalanceInputDTO {
  accountId: string
  balance: number
}

export interface UpdateBalanceOutputDTO {}

export class UpdateBalanceUseCase {
  constructor(private accountRepository: AccountRepository) {}

  async execute(
    input: UpdateBalanceInputDTO,
  ): Promise<UpdateBalanceOutputDTO | null> {
    const account = await this.accountRepository.findById(input.accountId)
    if (!account) {
      throw new Error('account not found')
    }

    account.balance = input.balance

    await this.accountRepository.updateBalance(account)

    return null
  }
}
