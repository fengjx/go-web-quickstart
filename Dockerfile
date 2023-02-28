FROM golang:1.20.1-alpine3.17 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -mod=readonly -v -o web-app ./cmd/main.go


FROM alpine:3.17
RUN apk --no-cache add ca-certificates bash

ENV LOG_DIR=/var/log/web
RUN mkdir -p ${LOG_DIR}

ENV APP_NAME=web-app
ENV WORK_DIR=/app

WORKDIR ${WORK_DIR}

COPY --from=build /app/${APP_NAME} .
COPY --from=build /app/build/*.sh .
COPY --from=build /app/configs ./configs
RUN ls -la

EXPOSE 8080
ENTRYPOINT ["sh", "-c", "./entrypoint.sh"]
