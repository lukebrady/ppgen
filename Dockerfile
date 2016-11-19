FROM ubuntu:latest
MAINTAINER Luke Brady

RUN apt-get update
RUN apt-get install -y golang git
RUN git clone

RUN go build /ppgen .
RUN ./ppgen.exe

EXPOSE 8080