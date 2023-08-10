import { Server } from 'socket.io'
import { ITelemetry, createProducer } from './rabbit-service'

// create socket to receive data
const io = new Server(4000)
console.log('Socket created at port 4000')

// start up the producer to send messages
const producer = createProducer('amqp://localhost', 'messaging')
console.log('Producer created to produce in amqp://localhost on queue "messaging"')

io.on('connection', (socket) => {
    socket.on('telemetry', (arg: string) => {
        const data: ITelemetry = JSON.parse(arg)

        console.log(`Received data via socket: ${data.speed}`)

        producer(data)

        socket.emit('ack', 'success')
    })

})