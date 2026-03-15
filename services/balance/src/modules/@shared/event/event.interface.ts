export default interface EventInterface<Payload = unknown> {
  getName(): string
  getDateTime(): Date
  getPayload(): Payload
  setPayload(payload: Payload): void
}
