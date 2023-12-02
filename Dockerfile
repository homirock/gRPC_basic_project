FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go build -o test1
EXPOSE 8084
CMD ["./test1"]