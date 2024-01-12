FROM node:20.6.0-alpine

WORKDIR /home

COPY package*.json ./
RUN npm i
COPY . .

CMD ["npm", "run", "dev"]