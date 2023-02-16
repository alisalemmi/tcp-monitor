#
# Build
#
FROM golang:1.20.1-alpine3.17 as build

WORKDIR /app/tcp-monitor

# download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build
COPY . .
RUN go build -o ./tcp-monitor

#
# Deploy
#
FROM alpine:3.17

WORKDIR /app/tcp-monitor

ENV PORT=3000
ENV SNMP_PATH=/host/net/snmp

EXPOSE ${PORT}

# copy builded app
COPY --from=build /app/tcp-monitor/tcp-monitor .

# run
CMD [ "./tcp-monitor" ]
