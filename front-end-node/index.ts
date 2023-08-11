import { io } from 'socket.io-client'

const socket = io('ws://localhost:4000')

setInterval(() => {
    const sendData = {
        data: Math.floor(Math.random() * 100),
        time: Date.now()
    }
    socket.emit('telemetry', JSON.stringify(sendData))
    console.log(`Sent data at ${Date.now()}`)

    socket.on('ack', (data) => {
        const latestMessage = `Received ack at ${Date.now()}. TAT = ${Date.now() - sendData.time}`
        console.clear()
        console.log(latestMessage)
    })
}, 3000)