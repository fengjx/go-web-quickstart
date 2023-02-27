FROM golang:1.20.1-alpine3.17 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -mod=readonly -v -o web-app ./cmd/main.go


FROM alpine:3.17
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/web

WORKDIR /app
COPY --from=build /app/web-app .
COPY --from=build /app/build/entrypoint.sh .
COPY --from=build /app/configs .
RUN ls -la

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
