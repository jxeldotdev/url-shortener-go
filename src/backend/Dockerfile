FROM golang:1.20-alpine AS build

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY . /app
RUN go build -o server ./cmd/server/main.go


FROM alpine:3.17.3

RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=build /app/server /usr/local/bin/
EXPOSE 8080
CMD ["/usr/local/bin/server"]
