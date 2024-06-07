FROM golang::1.22.3
WORKDIR app/
COPY . .
RUN go tidy
RUN go pkg/migrate/migrate.go
CMD ["go", "build", "-o", "bin", "."]
EXPOSE :8080
ENV DB='host=localhost user=postgres password=your_password dbname=name port=5432 sslmode=disable'
ENTRYPOINT ["app/bin"]

