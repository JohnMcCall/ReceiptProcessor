# syntax=docker/dockerfile:1

FROM golang:1.21.0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /receipt-processor
EXPOSE 8080
CMD ["/receipt-processor"]