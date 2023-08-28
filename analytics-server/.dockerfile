FROM alpine:latest

RUN mkdir /app

COPY /analyticsServerApp /app/analyticsServerApp

COPY ./connect-bundle-mqtt-storage-test.yml /connect-bundle-mqtt-storage-test.yml

CMD [ "/app/analyticsServerApp" ]