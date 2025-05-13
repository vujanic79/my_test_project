# ------------------------------
# Base image
# ------------------------------
FROM golang:1.24-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ------------------------------
# Development image
# ------------------------------
FROM base AS development

WORKDIR /app/cmd/web

RUN go build -o /app/bin/golang-react-todo-app

WORKDIR /app

EXPOSE 8000

# ------------------------------
# Builder image for production
# ------------------------------
FROM base AS builder

WORKDIR /app/cmd/web

RUN CGO_ENABLED=0 go build -o /app/bin/golang-react-todo-app

# ------------------------------
# Final production image
# ------------------------------
FROM scratch AS production

WORKDIR /app

COPY --from=builder /app/bin/golang-react-todo-app ./bin/golang-react-todo-app
COPY pkg/repository/data/ ./pkg/repository/data/

EXPOSE 8000