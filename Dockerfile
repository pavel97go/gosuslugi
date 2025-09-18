# --- build stage ---
FROM golang:1.25-bookworm AS builder
WORKDIR /app

# модули — для кеша
COPY go.mod go.sum ./
RUN go mod download

# код
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api ./cmd/api

# --- runtime stage ---
FROM gcr.io/distroless/base-debian12
WORKDIR /app

# конфиг положим внутрь образа под именем config.yaml
COPY config/config.docker.yaml /app/config/config.yaml
COPY --from=builder /api /app/api

EXPOSE 8080
ENTRYPOINT ["/app/api"]
