# Frontend build
FROM node:18-buster as nodebuild

WORKDIR /app

COPY ./frontend ./frontend
COPY ./tools/install-yarn.sh ./tools/install-yarn.sh
COPY ./tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY Makefile .

RUN make frontend

# Backend build
FROM golang:1.20-buster as gobuild

ENV GOPRIVATE=gitlab.com/datalift/*

WORKDIR /app

COPY ./backend ./backend
COPY ./tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY Makefile .
COPY --from=nodebuild /app/frontend/build ./frontend/build

RUN make backend-with-assets

# Copy binary to final image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=gobuild /app/build/server /app
COPY ./backend/datalift-config.yaml /app

EXPOSE 8080
CMD ["/app/server"]
