FROM golang:1.20-alpine as builder

ADD . /aki
WORKDIR /aki

ARG VERSION

RUN apk update
RUN apk add --no-cache git gcc musl-dev
RUN go build -ldflags="-X 'main.version=${VERSION}-$(git rev-parse --short HEAD)'" -o /tmp/aki ./

FROM jerray/goose:2.7.0-rc3 as goose

FROM alpine:3.15

ENV MIGRATIONS_PATH=/var/lib/aki/migrations

COPY --from=builder /tmp/aki /usr/bin/aki
COPY --from=builder /aki/config.yaml /var/lib/aki/config.yaml
COPY --from=goose /bin/goose /usr/bin/goose

ENV DOCS_PATH=/var/lib/aki/docs
ENV MIGRATIONS_PATH=/var/lib/aki/migrations
ENV HTTP_PORT=8001
ENV CONFIG_YAML=/var/lib/aki/config.yaml

COPY ./docs $DOCS_PATH
COPY ./migrations $MIGRATIONS_PATH

RUN apk apk update
RUN apk add --no-cache git

RUN echo '#! /bin/sh' > /usr/bin/entrypoint.sh
RUN echo 'goose -dir $MIGRATIONS_PATH postgres "$PG_DSN" up &&' >> /usr/bin/entrypoint.sh
RUN echo 'aki' >> /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh

EXPOSE 8000
ENTRYPOINT ["/usr/bin/entrypoint.sh"]