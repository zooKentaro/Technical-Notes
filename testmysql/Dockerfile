FROM golang:1.21.4-alpine3.18 AS builder
COPY /mysql/my.cnf /etc/my.cnf

RUN echo 'user=mysql' >> /etc/my.cnf