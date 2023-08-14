import { Kafka, KafkaMessage, Partitioners } from "kafkajs";

const kafka = new Kafka({
    brokers: ['pkc-41p56.asia-south1.gcp.confluent.cloud:9092'],
    sasl: {
        mechanism: 'plain',
        username: 'UKIAIS55B4VDG7RH',
        password: 'nws9Ep4BSpTrIlukaFNi21wg2O9/QkYQnKQ8B2hduu+8tbfwpiphC2x9kz/t78Mw'
    },
    ssl: true,
})

export interface IMessage {
    time: Date
    data: number
}

export class KafkaService {
    static async produce(data: IMessage, topic: string) {
        const producer = kafka.producer({ createPartitioner: Partitioners.LegacyPartitioner })
        await producer.connect()
        await producer.send({
            topic,
            messages: [
                { value: `${JSON.stringify(data)}` }
            ]
        })

        await producer.disconnect()
    }

    static async consume(group: string, topics: string[]) {
        const consumer = kafka.consumer({ groupId: group })

        consumer.on(consumer.events.CRASH, () => {
            return { message: 'crashed' }
        })

        const data: IMessage[] = []

        try {
            await consumer.connect()
            await consumer.subscribe({ topics, fromBeginning: true })

            await consumer.run({
                eachMessage: async ({ topic, partition, message }) => {
                    if (message.value?.toString())
                        data.push(JSON.parse(message.value.toString()))
                }
            })
        } catch (e) {
            console.log({ errors: e })
        }
        return { message: 'connected', data }
    }
}