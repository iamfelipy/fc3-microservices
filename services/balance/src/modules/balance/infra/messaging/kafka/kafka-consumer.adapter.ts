import {
  Message,
  MessageConsumerPort,
} from '@/modules/@shared/messaging/message-consumer.port'
import { Kafka, Consumer as KJSConsumer } from 'kafkajs'

interface KafkaConsumerConfig {
  clientId: string
  brokers: string[]
  groupId: string
  topics: string[]
}

export class KafkaConsumerAdapter implements MessageConsumerPort {
  private consumer: KJSConsumer

  constructor(private readonly config: KafkaConsumerConfig) {
    const kafka = new Kafka({
      clientId: config.clientId,
      brokers: config.brokers,
    })
    this.consumer = kafka.consumer({ groupId: config.groupId })
  }

  async init(onMessage: (message: Message) => Promise<void>): Promise<void> {
    await this.consumer.connect()
    await this.consumer.subscribe({
      topics: this.config.topics,
      fromBeginning: false,
    })
    await this.consumer.run({
      eachMessage: async ({ topic, message }) => {
        await onMessage({
          topic,
          value: message.value?.toString() ?? null,
          key: message.key?.toString() ?? null,
        })
      },
    })
  }

  async disconnect(): Promise<void> {
    await this.consumer.disconnect()
  }
}
