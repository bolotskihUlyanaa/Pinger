FROM golang:alpine
WORKDIR .
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o pinger ./cmd
CMD ["./pinger"]