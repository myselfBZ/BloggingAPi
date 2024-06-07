FROM golang:1.22.3
WORKDIR /app
COPY . .
RUN go mod download
RUN go run pkg/migrate/migrate.go
RUN go build -o bin .
ENV DB='host=localhost user=postgres password=nothing dbname=blog port=5432 sslmode=disable'
EXPOSE 8080
ENTRYPOINT ["./bin"]

