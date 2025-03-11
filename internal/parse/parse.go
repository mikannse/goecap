package parse

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// PacketData 结构体用于存储解析后的数据包信息
type PacketData struct {
	Index     int
	Timestamp time.Time
	SourceIP  net.IP
	Payload   []byte
}

// ParsePcapFile 解析捕获的 pcap 文件并将解析结果发送到 chan
func ParsePcapFile(pcapFile string, index int, packetChan chan<- PacketData) {
	handle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		fmt.Printf("Failed to open pcap file %s: %v\n", pcapFile, err)
		return
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		timestamp := packet.Metadata().Timestamp

		// 提取源 IP
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		var sourceIP net.IP
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			sourceIP = ip.SrcIP
		}

		// 提取 payload
		payload := []byte{}
		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			payload = appLayer.Payload()
		}

		packetChan <- PacketData{
			Index:     index,
			Timestamp: timestamp,
			SourceIP:  sourceIP,
			Payload:   payload,
		}
	}
}
