FROM golang:1.24.4

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1

RUN go build -ldflags "-s -w" -o ./unlock_test.so -buildmode=plugin ./plugin/main.go

RUN go build -buildmode=plugin -gcflags="all=-N -l" -o ./unlock_test_debug.so ./plugin/main.go