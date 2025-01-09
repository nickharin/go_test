Для комплиляции необходимо использовать следующие параметры


## on macOS
``
CC=i686-linux-musl-gcc \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=386 \
    go build \
        -ldflags="-s -w -linkmode external -extldflags '-static'" \
        -o main main.go
``
## on iSH
``
GOMAXPROCS=1 ./ish-go-demo
``