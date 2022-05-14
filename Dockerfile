FROM golang:1.18.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -tags netgo -o /app/server ./cmd/server/

FROM scratch
WORKDIR /app
COPY --from=builder /app/server ./
ENV PORT=8000
ENV DB_URL="postgres://localhost:5432/emp-demo?sslmode=disable"
CMD [ "/app/server" ]
