FROM alpine:latest

RUN mkdir /app

COPY /storageServerApp /app/storageServerApp

COPY ./connect-bundle-mqtt-storage-test.yml /connect-bundle-mqtt-storage-test.yml

CMD [ "/app/storageServerApp" ]