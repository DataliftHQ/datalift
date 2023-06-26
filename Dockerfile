# Web build
FROM node:18-buster as nodebuild

WORKDIR /app

COPY ../web ./web
COPY ../tools/install-yarn.sh ./tools/install-yarn.sh
COPY ../tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY ../Makefile .

RUN make web

# Go build
FROM golang:1.20-buster as gobuild

WORKDIR /app

COPY . .
COPY --from=nodebuild /app/web/build ./web/build

#RUN make server-with-assets
RUN make build

# Copy binary to final image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=gobuild /app/build/datalift /app
COPY ./datalift.yaml /app

EXPOSE 8080
ENTRYPOINT ["/app/datalift"]
