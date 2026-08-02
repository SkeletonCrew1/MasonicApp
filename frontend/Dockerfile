FROM node:alpine

WORKDIR /app

ARG VITE_POSTING_SERVICE_URL
ARG VITE_AUTH_SERVICE_URL

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

RUN npm install -g serve

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

CMD ["serve", "-s", "dist", "-l", "3000"]

