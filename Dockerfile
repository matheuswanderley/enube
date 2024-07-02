FROM golang:1.20 as builder


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN GOOS=linux GOARCH=amd64 go build -o main .


FROM alpine:latest


RUN apk --no-cache add ca-certificates


WORKDIR /app


COPY --from=builder /app/main .
COPY --from=builder /app/dados.xlsx .
COPY --from=builder /app/users.db .


RUN chmod +x ./main


EXPOSE 8080


CMD ["./main"]
