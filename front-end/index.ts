import { io } from 'socket.io-client'

const hostIP = 'localhost'
const port = 4000
const socket = io(`ws://${hostIP}:${port}`)

setInterval(() => {
    const data = {
        time: Date.now(),
        speed: Math.floor(Math.random() * 50)
    }
    socket.emit('telemetry', JSON.stringify(data))
    console.log('Emitted telemetry')
}, 3000)

socket.on('ack', (arg: string) => {
    console.log(`Ack: ${arg}`)
})
