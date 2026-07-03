FROM golang:1.25-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X github.com/kota69py/go-practicum/cmd.version=$(git describe --tags --always --dirty 2>/dev/null || echo dev)" -o /go-practicum .

FROM scratch

COPY --from=build /go-practicum /go-practicum
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go-practicum"]
