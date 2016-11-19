FROM ubuntu:latest
MAINTAINER Luke Brady

RUN apt-get update
RUN apt-get install -y golang git
RUN git clone https://github.com/lukebrains/ppgen

RUN go build .
RUN ./ppgen.exe

EXPOSE 8080