all: build

build: clean
	CGO_ENABLED=0;GOOS="linux" go build -ldflags "-s -w" -o docker-proxy

win: clean
	CGO_ENABLED=0;GOOS="windows" go build -ldflags "-s -w" -o docker-proxy.exe

mac: clean
	CGO_ENABLED=0;GOOS="darwin" go build -ldflags "-s -w" -o docker-proxy

clean:
	rm -rf docker-proxy docker-proxy.exe

.PHONY: all build win mac clean
