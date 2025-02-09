FROM golang:alpine
WORKDIR .
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd
CMD ["./server"]
