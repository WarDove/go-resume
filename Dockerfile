FROM golang:1.17-alpine as BUILD
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o main

FROM alpine:3.16 AS RUNTIME
WORKDIR /app
COPY --from=BUILD /build .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
