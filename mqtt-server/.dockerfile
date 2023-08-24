FROM alpine:latest

RUN mkdir /app

COPY /mqttServerApp /app/mqttServerApp

CMD [ "/app/mqttServerApp" ]