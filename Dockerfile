FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o rinha_backend_go .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/rinha_backend_go .
COPY .env .
EXPOSE 8080
CMD ["./rinha_backend_go"]
