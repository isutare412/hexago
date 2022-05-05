FROM golang:1.18.1-bullseye as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
      go build \
        -ldflags='-w -s' \
        -o server \
        ./cmd/...

FROM gcr.io/distroless/static-debian11

WORKDIR /app
COPY --from=builder /build/server ./
COPY configs configs
ENTRYPOINT [ "./server" ]
CMD [ "-config", "configs/local/config.yaml" ]
