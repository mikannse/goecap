package main

import (
	"os"
	"os/signal"
	"syscall"

	"ecapGrpc/internal/capture"
	"ecapGrpc/internal/parse"
	"ecapGrpc/internal/process"
)

func main() {
	// 创建一个通道来接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个通道来接收解析后的数据包
	packetChan := make(chan parse.PacketData)

	// 启动一个 goroutine 来处理解析后的数据包
	go process.ProcessPackets(packetChan)

	// 启动捕获进程
	capture.StartCapture(packetChan, signalChan)
}
