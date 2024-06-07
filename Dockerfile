FROM golang:1.22
RUN apt-get update -y
RUN apt-get install -y tzdata

# timezone env with default
ENV TZ=Asia/Singapore

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD [ "/app/vs133api/main" ]