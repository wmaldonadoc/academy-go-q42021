# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /api-wzln
ARG PORT
RUN echo "Dockerfile building in port $PORT"
EXPOSE ${PORT}

CMD [ "/api-wzln" ]