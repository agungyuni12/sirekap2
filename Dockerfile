FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o sirekap-app .

FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/sirekap-app .

# Copy static assets and templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/template ./template
RUN mkdir -p /app/static/uploads /app/output

# Set environment variables
ENV SERVER_ADDR=0.0.0.0:8080

EXPOSE 8080

CMD ["./sirekap-app"]
