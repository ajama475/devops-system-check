FROM golang:1.21
WORKDIR /app
COPY system_check.go .
RUN go build -o system_check system_check.go
CMD ["./system_check"]
