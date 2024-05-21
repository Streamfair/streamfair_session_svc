# Build Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_session_svc
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o session_svc main.go

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_session_svc

# Copy the binary from the build stage
COPY --from=build /streamfair_session_svc/session_svc .

COPY sh ./sh
COPY db/migration ./db/migration

EXPOSE 8083
EXPOSE 9093

CMD [ "/streamfair_session_svc/session_svc" ]
ENTRYPOINT [ "/streamfair_session_svc/start.sh" ]

RUN apk add --no-cache bash curl