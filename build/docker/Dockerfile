FROM golang:1.20
WORKDIR /app
COPY . .
RUN go build -o weather-app ./server/main.go
CMD ["./weather-app"]
