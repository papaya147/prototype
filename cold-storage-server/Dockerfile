FROM --platform=arm64 node:16-alpine

WORKDIR /usr/app

COPY ./package.json ./

RUN npm install

COPY ./ ./

CMD [ "npm", "start" ]