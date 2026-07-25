FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app

RUN apt-get update && \
    apt-get install -y postgresql postgresql-contrib && \
    rm -rf /var/lib/apt/lists/*

COPY init.sql .
COPY start.sh .
COPY users_init.sql .

RUN chmod +x start.sh

EXPOSE 5432

CMD ["./start.sh"]