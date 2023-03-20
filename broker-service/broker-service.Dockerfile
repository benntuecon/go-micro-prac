# base go

FROM golang:1.18-alpine as builder


RUN mkdir /app

COPY . /app

WORKDIR 

