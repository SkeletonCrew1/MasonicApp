FROM node:alpine

WORKDIR /app

ARG VITE_POSTING_SERVICE_URL
ARG VITE_AUTH_SERVICE_URL
ARG VITE_BACKEND_DJANGO_URL

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN npm install -g serve



CMD [ "npm", "run", "dev", "--", "--host" ]