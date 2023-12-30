FROM golang:alpine

WORKDIR /dnbbotapi

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG DNBBOT_MONGO_PASS
ARG DNBBOT_ENV

ENV DNBBOT_MONGO_PASS ${DNBBOT_MONGO_PASS}
ENV DNBBOT_ENV ${DNBBOT_ENV}

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api

EXPOSE 8080
CMD ["/dnbbotapi/bin/api"]