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
const socket_io_1 = require("socket.io");
const kafka_service_1 = require("./kafka-service");
// create socket to receive data
const io = new socket_io_1.Server(4000);
console.log('Socket created at port 4000');
io.on('connection', (socket) => {
    socket.on('telemetry', (arg) => __awaiter(void 0, void 0, void 0, function* () {
        const data = JSON.parse(arg);
        console.log(`Received data via socket: ${data.data}`);
        // produce data to Kafka 
        yield kafka_service_1.KafkaService.produce(data, 'messages');
        console.log('Data produced to Kafka queue');
        socket.emit('ack', 'success');
    }));
});
