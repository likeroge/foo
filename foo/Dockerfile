# Этап сборки
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем статический бинарник
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o myapp .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o myapp ./cmd/api/main.go


# Финальный этап - минимальный образ
FROM alpine:3.18

# Устанавливаем зависимости времени выполнения если нужны
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из этапа сборки
COPY --from=builder /app/myapp .

RUN mkdir -p /root/data

# RUN mkdir -p /root/data && \
#     chown -R myuser:appgroup /root


# Создаем непривилегированного пользователя
RUN adduser -D myuser
USER myuser

EXPOSE 5000

CMD ["./myapp"]
