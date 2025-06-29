# Build binary
FROM golang:1.24.3-alpine3.22 AS builder

WORKDIR /app

COPY go.mod main.go ./

RUN go build -o main ./... 

# Run application
FROM alpine:3.22 AS runner

COPY --from=builder ./app/main .

CMD [ "./main" ]