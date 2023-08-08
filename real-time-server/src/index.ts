import { Server } from 'socket.io'
import { createProducer } from './rabbit-service'

const io = new Server(4000)

const producer = createProducer('amqp://localhost', 'messaging')

setInterval(() => {
    producer({ type: 'text', content: 'hi' })
    console.log('Message produced')
}, 3000)

