FROM golang:1.23
    
WORKDIR /app
    
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Экспонируем порт
EXPOSE 8080

# Запускаем приложение
CMD ["go", "run", "cmd/app/main.go"]
