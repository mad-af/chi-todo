FROM golang:1.18

WORKDIR /todo

RUN apt-get update

COPY . .

RUN go mod download \
 && go build -tags musl -o main ./bin/app

# Expose port
EXPOSE 9000

CMD ["go", "run", "./bin/app/main.go"]