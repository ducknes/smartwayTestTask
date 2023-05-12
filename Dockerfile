FROM golang:latest

COPY ./ ./
RUN go build -o swtt .
CMD ["./swtt"]