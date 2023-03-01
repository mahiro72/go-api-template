FROM golang:1.19 AS dev

WORKDIR /go/src/github.com/mahiro72/go-api-template

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@v1.29.0

COPY docker/dev/api/.air.toml /tmpair/.air.toml

CMD ["air", "-c", "/tmpair/.air.toml"]
