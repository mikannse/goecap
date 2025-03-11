package process

import (
	"ecapGrpc/internal/parse"
	"fmt"
)

// ProcessPackets 处理从 chan 接收到的数据包
func ProcessPackets(packetChan <-chan parse.PacketData) {
	for packetData := range packetChan {
		fmt.Printf("Processed packet from file %d.pcapng\n", packetData.Index)
		fmt.Printf("Timestamp: %v\n", packetData.Timestamp)
		fmt.Printf("Source IP: %v\n", packetData.SourceIP)
		fmt.Printf("Payload: %v\n", packetData.Payload)
		// 在这里可以对 packetData 进行进一步处理
	}
}
