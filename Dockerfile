FROM golang:1.24-alpine AS builder

WORKDIR /api

COPY go.mod go.sum ./

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /api/monitorx ./cmd/server/main.go

FROM alpine AS runtime

WORKDIR /api

COPY --from=builder /api/monitorx .

COPY --from=builder /api/ .

EXPOSE 5758

ENTRYPOINT [ "/api/monitorx" ]
