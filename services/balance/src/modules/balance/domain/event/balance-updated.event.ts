import EventInterface from '@/modules/@shared/event/event.interface'

export default class BalanceUpdatedEvent<T = unknown>
  implements EventInterface<T>
{
  private name: string
  private payload!: T
  private dateTime: Date

  constructor() {
    this.name = 'BalanceUpdated'
    this.dateTime = new Date()
  }

  getName(): string {
    return this.name
  }

  getPayload(): T {
    return this.payload
  }

  setPayload(payload: T): void {
    this.payload = payload
  }

  getDateTime(): Date {
    return this.dateTime
  }
}
