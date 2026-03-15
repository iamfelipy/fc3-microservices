import { AccountRepository } from "@/modules/balance/domain/repository/account.repository";
import { PrismaService } from "../prisma.service";
import { Account } from "@/modules/balance/domain/entity/account";

export class PrismaAccountRepository implements AccountRepository {
  constructor(private prisma: PrismaService) {}

  async findById(id: string): Promise<Account | null> {
    const accountData = await this.prisma.account.findUnique({
      where: { id },
    });

    if (!accountData) {
      return null;
    }

    return new Account(
      accountData.id,
      accountData.clientId,
      accountData.balance.toNumber(),
      accountData.createdAt,
    );
  }

  async updateBalance(account: Account): Promise<void> {
    await this.prisma.account.update({
      where: { id: account.id },
      data: {
        balance: account.balance,
      }
    });
  }
}