FROM golang:1.26.4-alpine AS builder
WORKDIR /services/govno
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM gcr.io/distroless/static-debian12
WORKDIR /services/govno
COPY --from=builder /services/govno/app ./app
CMD [ "./app" ]