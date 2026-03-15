export interface Message {
  topic: string
  value: string | null
  key: string | null
  headers?: Record<string, string>
}

export interface MessageConsumerPort {
  init(onMessage: (message: Message) => Promise<void>): Promise<void>
  disconnect(): Promise<void>
}
