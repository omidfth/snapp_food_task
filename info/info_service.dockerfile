# base go image
FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app/cmd/api

RUN CGO_ENABLED=0 go build -o infoApp .

RUN chmod +x /app/cmd/api/infoApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/cmd/api/infoApp /app

COPY .env /app

WORKDIR /app

CMD [ "/app/infoApp","-env","/app/.env"]