import { Client } from "./client"

export class Account{
  private _id: string
  private _client: string
  private _balance: number
  private _createdAt: Date
  private _updatedAt?: Date | null

  constructor(
    id: string,
    client: string,
    balance: number,
    createdAt?: Date,
    updatedAt?: Date | null
  ) {
    this._id = id;
    this._client = client;
    this._balance = balance;
    this._createdAt = createdAt ?? new Date();
    this._updatedAt = updatedAt;
  }


  get id() {
    return this._id;
  }

  get client() {
    return this._client;
  }


  get balance() {
    return this._balance;
  }

  set balance(value: number) {
    this._balance = value;
    this.touch();
  }

  get createdAt() {
    return this._createdAt;
  }

  get updatedAt() {
    return this._updatedAt;
  }

  private touch() {
    this._updatedAt = new Date()
  }

  credit(amount: number) {
    this._balance = this._balance + amount;
    this.touch();
  }

  debit(amount: number) {
    this._balance = this._balance - amount;
    this.touch();
  }
}
