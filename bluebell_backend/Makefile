.PHONY: all build run gotool clean help

BINARY="bluebell"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/${BINARY}

run:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"