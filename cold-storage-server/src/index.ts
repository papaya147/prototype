import mongoose from 'mongoose'
import { KafkaService } from './kafka-service'
import express from 'express'

// create database connection
const connect = async () => {
    try {
        const user = 'asrivatsa6'
        const pass = 'SOMGR1ioxM2eqhiL'
        const db = `mongodb+srv://${user}:${pass}@cluster0.mvlochz.mongodb.net`
        const connectionString = `${db}/telemetry`
        await mongoose.connect(connectionString)
        console.log('Connected to MongoDb')
    } catch (err) {
        console.error(err)
    }
}
connect()

// start up the consumer to consume messages
const kafka_start = async () => {
    let status = await KafkaService.consume('storage', ['messages'])
    if (status.message === 'crashed')
        kafka_start()
    else
        console.log('Connected to Kafka Brokers')
}

kafka_start()

const app = express()

app.get('/health', (req, res) => {
    res.sendStatus(200)
})

app.listen(4000, () => {
    console.log('Cold storage server running on port 4000')
})