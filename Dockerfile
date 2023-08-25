FROM golang:alpine3.18 AS build

WORKDIR /app
COPY . /app/
RUN go build -o main


FROM alpine:3.18
COPY --from=build /app/main .
CMD [ "/main" ]
