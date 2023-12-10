FROM golang:alpine AS builder

WORKDIR /app

# dependencies cache 
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o capital-gains .


FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/capital-gains .
CMD ["./capital-gains"]