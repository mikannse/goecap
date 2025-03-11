package capture

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"ecapGrpc/internal/parse"
	"ecapGrpc/internal/utils"
)

// StartCapture 启动捕获进程
func StartCapture(packetChan chan<- parse.PacketData, signalChan chan os.Signal) {
	// 获取最新的文件名序号
	index := utils.GetLastIndex()

	for {
		// 定义当前文件名
		pcapFile := filepath.Join(utils.OutputDir, fmt.Sprintf("%d.pcapng", index))

		// 启动新的捕获进程
		currentCmd := exec.Command("sudo", "./ecapture", "tls", "-m", "pcapng", "-i", "ens33", "--pcapfile", pcapFile)
		if err := currentCmd.Start(); err != nil {
			fmt.Printf("Failed to start ecapture: %v\n", err)
			return
		}

		fmt.Printf("Starting capture for file %d.pcapng\n", index)

		// 等待60秒
		select {
		case <-time.After(60 * time.Second):
			// 终止上一个进程
			if currentCmd.Process != nil {
				if err := currentCmd.Process.Kill(); err != nil {
					fmt.Printf("Failed to kill process: %v\n", err)
				} else {
					fmt.Printf("Capture for file %d.pcapng completed and saved successfully\n", index)
				}
			}

			// 解析捕获的文件
			go parse.ParsePcapFile(pcapFile, index, packetChan)
		case <-signalChan:
			// 处理终止信号
			if currentCmd.Process != nil {
				if err := currentCmd.Process.Kill(); err != nil {
					fmt.Printf("Failed to kill process: %v\n", err)
				} else {
					fmt.Printf("Capture for file %d.pcapng completed and saved successfully\n", index)
				}
			}
			return
		}

		// 更新文件名序号
		index++
		utils.WriteLastIndex(index)
	}
}
