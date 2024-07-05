FROM dnsunlockcom/build-go-1.22.5:cn

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1

RUN go build -ldflags "-s -w" -o ./unlock_test.so -buildmode=plugin ./plugin/main.go