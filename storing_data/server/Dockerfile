FROM golang:1.20-alpine
WORKDIR /serv 
COPY . .
RUN go mod tidy
RUN go build -o server 
EXPOSE 8080
CMD ["./server"] 
