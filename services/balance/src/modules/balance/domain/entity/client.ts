import { Account } from "./account"

export class Client{
  private _id: string
  private _name: string
  private _email: string
  private _accounts: Account[]
  private _createdAt: Date
  private _updatedAt?: Date | null

  constructor(
    id: string,
    name: string,
    email: string,
    accounts: Account[],
    createdAt?: Date,
    updatedAt?: Date | null
  ) {
    this._id = id;
    this._name = name;
    this._email = email;
    this._accounts = accounts;
    this._createdAt = createdAt ?? new Date();
    this._updatedAt = updatedAt;
  }

  get id(): string {
    return this._id;
  }

  get name(): string {
    return this._name;
  }

  get email(): string {
    return this._email;
  }

  get accounts(): Account[] {
    return this._accounts;
  }

  set accounts(value: Account[]) {
    this._accounts = value;
    this.touch();
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  get updatedAt() {
    return this._updatedAt;
  }

  private touch() {
    this._updatedAt = new Date()
  }

  validate(): void {
    if (!this._name || this._name.trim() === "") {
      throw new Error("name is required");
    }
    if (!this._email || this._email.trim() === "") {
      throw new Error("email is required");
    }
  }

  update(name: string, email: string): void {
    this._name = name;
    this._email = email;
    this.touch();
    this.validate();
  }

  addAccount(account: Account): void {
    if (account.client.id !== this._id) {
      throw new Error("account does not belong to client");
    }
    this._accounts.push(account);
    this.touch();
  }
}
