# Build Stage
FROM golang:1.22.0-alpine3.19 AS build
WORKDIR /streamfair_session_svc
COPY . .
RUN go build -o session_svc main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.19
WORKDIR /streamfair_session_svc

# Copy the binary from the build stage
COPY --from=build /streamfair_session_svc/session_svc .
# Copy the downloaded migration binary from the build stage
COPY --from=build /streamfair_session_svc/migrate ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8083
EXPOSE 9093

CMD [ "/streamfair_session_svc/session_svc" ]
ENTRYPOINT [ "/streamfair_session_svc/start.sh" ]