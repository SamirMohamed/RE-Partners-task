FROM golang:alpine3.18

WORKDIR /app

COPY main.go /app/

CMD [ "go", "run", "main.go" ]
