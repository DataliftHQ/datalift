# Web build
FROM node:18-buster as nodebuild

WORKDIR /app

COPY ../ui ./ui
COPY ../tools/install-yarn.sh ./tools/install-yarn.sh
COPY ../tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY ../Makefile .

RUN make web

# Server build
FROM golang:1.20-buster as gobuild

WORKDIR /app

COPY ./api ./api
COPY ./server ./server
COPY ../tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY ../Makefile .
COPY --from=nodebuild /app/ui/build ./ui/build

RUN make server-with-assets

# Copy binary to final image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=gobuild /app/build/server /app
COPY ./config/datalift-config.yaml /app

EXPOSE 8080
CMD ["/app/server"]
