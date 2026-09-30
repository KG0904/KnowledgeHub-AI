FROM golang:1.26.5

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . .

RUN go build -o go-gateway ./engine

EXPOSE 8080

CMD ["./go-gateway"]