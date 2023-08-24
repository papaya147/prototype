FROM alpine:latest

RUN mkdir /app

COPY /mqttClientApp /app/mqttClientApp

CMD [ "/app/mqttClientApp" ]