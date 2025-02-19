FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go env -w GOPROXY="https://proxy.golang.org,direct"
RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 go build -o nimbus-backend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/nimbus-backend .

EXPOSE 8080

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

CMD ["./nimbus-backend"]
