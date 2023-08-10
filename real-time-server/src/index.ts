import { Server } from 'socket.io'
import { IMessage, KafkaService } from './kafka-service'

// create socket to receive data
const io = new Server(4000)
console.log('Socket created at port 4000')

io.on('connection', (socket) => {
    socket.on('telemetry', async (arg: string) => {
        const data: IMessage = JSON.parse(arg)

        console.log(`Received data via socket: ${data.data}`)

        // produce data to Kafka 
        await KafkaService.produce(data, 'messages')
        console.log('Data produced to Kafka queue')

        socket.emit('ack', 'success')
    })
})