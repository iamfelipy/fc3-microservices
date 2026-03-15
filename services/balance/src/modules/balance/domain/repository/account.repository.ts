import { Account } from '@/modules/balance/domain/entity/account'

export interface AccountRepository {
  findById(id: string): Promise<Account | null>
  updateBalance(account: Account): Promise<void>
}
