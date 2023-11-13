FROM golang:1.19.4

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

EXPOSE 8080

CMD ["./field-team-management"]
