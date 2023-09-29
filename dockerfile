FROM golang:1.21.1

WORKDIR /app
copy . .
RUN go get
RUN go build -o bin .
ENTRYPOINT [ "/app/bin" ]
