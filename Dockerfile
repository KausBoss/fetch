FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o receiptprocessor .

CMD ["./receiptprocessor"]
