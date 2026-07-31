FROM node:alpine

WORKDIR /app

ARG VITE_POSTING_SERVICE_URL

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build
RUN npm install -g serve

COPY docker-entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]