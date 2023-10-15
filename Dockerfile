FROM golang:1.21.3 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch AS final

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./main"]