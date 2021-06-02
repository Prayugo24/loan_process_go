FROM golang:1.15.12

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 9008

RUN go build

CMD ["./tunaiku_tes"]