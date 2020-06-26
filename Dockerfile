FROM golang:1.13-alpine AS server-builder

RUN apk --no-cache add make
WORKDIR /go/src/github.com/theoptz/grpcchat

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN make server

FROM scratch

WORKDIR /go/src/github.com/theoptz/grpcchat
COPY --from=server-builder /go/src/github.com/theoptz/grpcchat/cmd/server/server ./cmd/server/server
EXPOSE 8080

CMD ["./cmd/server/server", "-a=0.0.0.0:8080", "-d=true"]