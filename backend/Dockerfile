FROM golang:1.22-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o telemetry-api .

FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=build /app/telemetry-api /app/telemetry-api

EXPOSE 8080
ENV PORT=8080
ENV DB_PATH=/data/telemetry.db

CMD ["/app/telemetry-api"]
