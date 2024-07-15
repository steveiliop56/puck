# ---- BUILDER ----
FROM golang:1.23-rc as builder

WORKDIR /puck

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./main.go ./
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build .

# ---- RUNNER ----
FROM busybox:1.36.1 as runner

WORKDIR /

COPY --from=builder /puck/puck ./puck

ENTRYPOINT ['/puck']
