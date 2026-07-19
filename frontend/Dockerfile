FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache \
    nodejs \
    npm

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]