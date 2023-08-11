import { Kafka, KafkaMessage, Partitioners } from "kafkajs";
import { Telemetry } from "./telemetry.model";

const kafka = new Kafka({
    brokers: ['adequate-anchovy-12371-eu1-kafka.upstash.io:9092'],
    sasl: {
        mechanism: 'scram-sha-256',
        username: 'YWRlcXVhdGUtYW5jaG92eS0xMjM3MSQwl0sMEaYfw3sxnQv1yUfPExdOBr7n-nI',
        password: 'a74f1e7b4d864ba4ae4b8ffca7a8b88c'
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

        try {
            await consumer.connect()
            await consumer.subscribe({ topics, fromBeginning: true })

            await consumer.run({
                eachMessage: async ({ topic, partition, message }) => {
                    if (message.value?.toString())
                        switch (topic) {
                            case 'messages':
                                this.handleData(JSON.parse(message.value?.toString()))
                                break
                        }
                }
            })
        } catch (e) {
            console.log({ errors: e })
        }
        return { message: 'connected' }
    }

    static async handleData(message: IMessage) {
        const telemetry = Telemetry.build(message)
    }
}