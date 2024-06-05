FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD [ "/app/vs133api/main" ]