# ---------- BUILD STAGE ----------
FROM golang:1.25.7-alpine AS builder

WORKDIR /app

# install git (нужно для go mod download)
RUN apk add --no-cache git

# copy go mod files first (кэшируется)
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bridge-server main.go


# ---------- RUNTIME STAGE ----------
FROM alpine:3.19

WORKDIR /app

# add ca-certificates for HTTPS (Alchemy, Infura, etc)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bridge-server .

COPY .env .

EXPOSE 8080

CMD ["./bridge-server"]