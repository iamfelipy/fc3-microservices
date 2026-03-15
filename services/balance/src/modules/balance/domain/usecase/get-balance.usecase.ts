import { AccountRepository } from '@/modules/balance/domain/repository/account.repository'

export interface GetBalanceInputDTO {
  accountId: string
}

export interface GetBalanceOutputDTO {
  accountId: string
  balance: number
}

export class GetBalanceUseCase {
  constructor(private accountRepository: AccountRepository) {}

  async execute(input: GetBalanceInputDTO): Promise<GetBalanceOutputDTO> {
    const account = await this.accountRepository.findById(input.accountId)
    if (!account) {
      throw new Error('account not found')
    }

    return {
      accountId: account.id,
      balance: account.balance,
    }
  }
}
