FROM ubuntu:latest

WORKDIR /go

COPY sender-linux /go

RUN chmod +x sender-linux
RUN apt-get update && \
    apt-get install tcpdump -y

# CMD ["./sender"]

