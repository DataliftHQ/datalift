# Worker build
FROM golang:1.20-buster as gobuild

WORKDIR /app

COPY ./worker ./worker
COPY ./tools/preflight-checks.sh ./tools/preflight-checks.sh
COPY ./Makefile .

RUN make worker

# Copy binary to final image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=gobuild /app/build/worker /app

CMD ["/app/worker"]
