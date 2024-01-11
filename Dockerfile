# Используем образ Golang
FROM golang:1.21-alpine

# Выводим информацию о версии Go
RUN go version

# Устанавливаем переменную окружения GOPATH
ENV GOPATH=/

# Копируем все файлы проекта в текущую директорию образа
COPY ./ .

# Загружаем зависимости модулей Go
RUN go mod download

# Собираем проект в бинарный файл my_hata
RUN go build -o my_hata ./main.go

# Указываем команду для запуска приложения
CMD ["./my_hata"]
