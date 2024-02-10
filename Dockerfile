# Build Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_session_service
COPY . .
RUN go build -o streamfair_session_service main.go

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_session_service
COPY --from=build /streamfair_session_service/streamfair_session_service .
COPY app.env .

EXPOSE 8083
EXPOSE 9093

CMD ["./streamfair_session_service"]