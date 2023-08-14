import { Server } from 'socket.io'
import { IMessage, KafkaService } from './kafka-service'
import express from 'express'

// create socket to receive data
const io = new Server(4000)
console.log('Socket created at port 4000')

const app = express()

io.on('connection', (socket) => {
    socket.on('telemetry', async (arg: string) => {
        const data: IMessage = JSON.parse(arg)

        console.log(`Received data via socket: data = ${data.data}, time = ${data.time}`)

        // produce data to Kafka 
        KafkaService.produce(data, 'messages')

        socket.emit('ack', JSON.stringify(data))
    })
})

app.get('/health', (req, res) => {
    res.sendStatus(200)
})

app.listen(4001, () => {
    console.log('Real time server running on port 4000')
})