FROM golang:1.22.3
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bin .
ENV DB='host=db user=myuser password=mysecretpassword dbname=blog port=5432 sslmode=disable'
EXPOSE 8080
CMD ["go", "run", "app/migrate/migrate.go"]
ENTRYPOINT ["./bin"]
