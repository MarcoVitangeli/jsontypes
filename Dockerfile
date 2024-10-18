FROM golang:1.23 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . /app
RUN templ generate cmd/web
RUN go build ./cmd/web/

FROM ubuntu:jammy AS release-stage
WORKDIR /
COPY --from=build-stage /app/web /web
EXPOSE 3000
ENTRYPOINT ["/web"]