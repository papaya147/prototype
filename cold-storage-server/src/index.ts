import { createConsumer } from './rabbit-service'

const consumer = createConsumer('amqp://localhost', 'messaging')

consumer()