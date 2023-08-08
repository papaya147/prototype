import amqp, { Connection, Channel, Message } from 'amqplib/callback_api';

export const createProducer = (amqpUrl: string, queueName: string) => {
    let ch: Channel

    // check connection to url provided
    amqp.connect(amqpUrl, (errorConnect: Error, connection: Connection) => {
        if (errorConnect)
            throw errorConnect

        // create channel to pass messages
        connection.createChannel((errorChannel: Error, channel: Channel) => {
            if (errorChannel)
                throw errorChannel

            ch = channel
        })
    })

    // return method to produce messages on a specified queue
    return (message: ITelemetry) => {
        console.log('Producing message to RabbitMQ')
        ch.sendToQueue(queueName, Buffer.from(JSON.stringify(message)))
    }
}

export interface ITelemetry {
    time: Date
    speed: number
}

export const createConsumer = (amqpUrl: string, queueName: string) => {
    console.log('Connecting to RabbitMQ...')

    // return method to consume messages on specified queue
    return () => {
        // check connection to url provided
        amqp.connect(amqpUrl, (errorConnect: Error, connection: Connection) => {
            if (errorConnect)
                throw errorConnect

            // create channel to receive messages
            connection.createChannel((errorChannel: Error, channel: Channel) => {
                if (errorChannel)
                    throw errorChannel

                console.log('Connected to RabbitMQ')

                // check if the specified queue exists, if not create one
                channel.assertQueue(queueName, { durable: true })

                // consume messages from the channel
                channel.consume(queueName, (msg: Message | null) => {
                    if (msg) {
                        // parse message content
                        const parsed: ITelemetry = JSON.parse(msg.content.toString())

                        // do something with the data (put in DB, return it, append to array, etc.)
                        console.log(`Data received: ${parsed.time}, ${parsed.speed}`)
                    }
                }, { noAck: true })
            })
        })
    }
}