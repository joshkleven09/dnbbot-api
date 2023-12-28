FROM golang:alpine

WORKDIR /dnbbotapi

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
# todo env based config copying here
# todo secrets copying here
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api

EXPOSE 8080
CMD ["/dnbbotapi/bin/api"]