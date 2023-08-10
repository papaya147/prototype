"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.KafkaService = void 0;
const kafkajs_1 = require("kafkajs");
const kafka = new kafkajs_1.Kafka({
    brokers: ['adequate-anchovy-12371-eu1-kafka.upstash.io:9092'],
    sasl: {
        mechanism: 'scram-sha-256',
        username: 'YWRlcXVhdGUtYW5jaG92eS0xMjM3MSQwl0sMEaYfw3sxnQv1yUfPExdOBr7n-nI',
        password: 'a74f1e7b4d864ba4ae4b8ffca7a8b88c'
    },
    ssl: true,
});
class KafkaService {
    static produce(data, topic) {
        return __awaiter(this, void 0, void 0, function* () {
            const producer = kafka.producer({ createPartitioner: kafkajs_1.Partitioners.LegacyPartitioner });
            yield producer.connect();
            yield producer.send({
                topic,
                messages: [
                    { value: `${JSON.stringify(data)}` }
                ]
            });
            yield producer.disconnect();
        });
    }
    static consume(group, topics) {
        return __awaiter(this, void 0, void 0, function* () {
            const consumer = kafka.consumer({ groupId: group });
            consumer.on(consumer.events.CRASH, () => {
                return { message: 'crashed' };
            });
            const data = [];
            try {
                yield consumer.connect();
                yield consumer.subscribe({ topics, fromBeginning: true });
                yield consumer.run({
                    eachMessage: ({ topic, partition, message }) => __awaiter(this, void 0, void 0, function* () {
                        var _a;
                        if ((_a = message.value) === null || _a === void 0 ? void 0 : _a.toString())
                            data.push(JSON.parse(message.value.toString()));
                    })
                });
            }
            catch (e) {
                console.log({ errors: e });
            }
            return { message: 'connected', data };
        });
    }
}
exports.KafkaService = KafkaService;
