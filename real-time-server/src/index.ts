import { Server } from 'socket.io'
import { IMessage, KafkaService } from './kafka-service'
import fs from 'fs'
import https from 'https'

const options = {
    key: fs.readFileSync('privatekey.pem'),
    cert: fs.readFileSync('certificate.pem')
}

// create https server
const server = https.createServer(options)

// create socket to receive data
const io = new Server(server)
console.log('Socket created at port 4000')

io.on('connection', (socket) => {
    socket.on('telemetry', async (arg: string) => {
        const data: IMessage = JSON.parse(arg)

        console.log(`Received data via socket: data = ${data.data}, time = ${data.time}`)

        // produce data to Kafka 
        KafkaService.produce(data, 'messages')

        socket.emit('ack', JSON.stringify(data))
    })
})

server.listen(4000, () => {
    console.log('Server is running on port 4000');
})