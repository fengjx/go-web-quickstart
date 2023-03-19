FROM golang:1.20.1-alpine3.17 AS build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY build ./build
COPY cmd ./cmd
COPY configs ./configs
COPY internal ./internal
COPY pkg ./pkg
COPY Makefile ./
RUN CGO_ENABLED=0 go build -tags=jsoniter -mod=readonly -v -o app ./cmd/main.go


FROM alpine:3.17
RUN apk --no-cache add ca-certificates bash curl

ENV LOG_DIR=/var/log/web
RUN mkdir -p ${LOG_DIR}

ENV APP_NAME=app
ENV WORK_DIR=/app

WORKDIR ${WORK_DIR}

COPY --from=build /app/${APP_NAME} .
COPY --from=build /app/build/*.sh .
COPY --from=build /app/configs ./configs
RUN ls -la

EXPOSE 8080
ENTRYPOINT ["sh", "-c", "./entrypoint.sh"]
