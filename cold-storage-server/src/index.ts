import mongoose from 'mongoose'
import { createConsumer } from './rabbit-service'

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
const consumer = createConsumer('amqp://localhost', 'messaging')
consumer()
console.log('Consumer created to consume from amqp://localhost on queue "messaging"')
