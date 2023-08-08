import amqp, { Connection, Channel, Message } from 'amqplib/callback_api';

export const createProducer = (amqpUrl: string, queueName: string) => {
    let ch: Channel

    amqp.connect(amqpUrl, (errorConnect: Error, connection: Connection) => {
        if (errorConnect)
            throw errorConnect

        connection.createChannel((errorChannel: Error, channel: Channel) => {
            if (errorChannel)
                throw errorChannel

            ch = channel
        })
    })

    return (message: IMsg) => {
        console.log('Producing message to RabbitMQ')
        ch.sendToQueue(queueName, Buffer.from(JSON.stringify(message)))
    }
}

interface IMsg {
    type: string
    content: string
}

export const createConsumer = (amqpUrl: string, queueName: string) => {
    console.log('Connecting to RabbitMQ...')
    return () => {
        amqp.connect(amqpUrl, (errorConnect: Error, connection: Connection) => {
            if (errorConnect)
                throw errorConnect

            connection.createChannel((errorChannel: Error, channel: Channel) => {
                if (errorChannel)
                    throw errorChannel

                console.log('Connected to RabbitMQ')
                channel.assertQueue(queueName, { durable: true })
                channel.consume(queueName, (msg: Message | null) => {
                    if (msg) {
                        console.log(msg.content)
                        // const parsed: IMsg = JSON.parse(msg.content.toString())
                        // console.log(`Data received: ${parsed.content}`)
                    }
                }, { noAck: true })
            })
        })
    }
}