FROM golang:1.20-bookworm

ENV HOME /home/local/
WORKDIR /home/local/

RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y sqlite3
