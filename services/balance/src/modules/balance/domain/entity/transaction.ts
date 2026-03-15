import { Account } from "./account"

export class Transaction {
  private _id: string
  private _accountFrom: Account
  private _accountTo: Account
  private _amount: number
  private _createdAt: Date

  constructor(
    id: string,
    accountFrom: Account,
    accountTo: Account,
    amount: number,
    createdAt?: Date,
  ) {
    this._id = id;
    this._accountFrom = accountFrom;
    this._accountTo = accountTo;
    this._amount = amount;
    this._createdAt = createdAt ?? new Date();

    this.validate();
    this.commit();
  }

  get id() {
    return this._id;
  }

  get accountFrom() {
    return this._accountFrom;
  }

  get accountTo() {
    return this._accountTo;
  }

  get amount() {
    return this._amount;
  }

  get createdAt() {
    return this._createdAt;
  }

  private commit() {
    this._accountFrom.debit(this._amount);
    this._accountTo.credit(this._amount);
  }

  private validate() {
    if (this._amount <= 0) {
      throw new Error("amount must be greater than zero");
    }
    if (this._accountFrom.balance < this._amount) {
      throw new Error("insufficient funds");
    }
  }
}
