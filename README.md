# ecapGrpc 项目

## 项目简介

`ecapGrpc` 是一个基于 Go 语言开发的项目，主要用于通过 gRPC 提供对 `pcap` 文件解析的服务。它依赖于 `github.com/google/gopacket` 库来处理和解析网络数据包。

## 依赖项

- `github.com/google/gopacket`：用于解析 `pcap` 文件。

## 安装步骤

1. **安装 Go 语言环境**：
   确保你已经安装了 Go 语言环境。你可以从 [Go 官方网站](https://golang.org/dl/) 下载并安装最新版本的 Go。

2. **安装 `libpcap` 开发库**：
   `gopacket` 依赖于 `libpcap` 库。根据你的操作系统安装相应的开发库。

   - **Debian/Ubuntu**:
     ```bash
     sudo apt-get update
     sudo apt-get install -y libpcap-dev
     ```

   - **CentOS/RHEL**:
     ```bash
     sudo yum install -y libpcap-devel
     ```

   - **Fedora**:
     ```bash
     sudo dnf install -y libpcap-devel
     ```

   - **macOS** (使用 Homebrew):
     ```bash
     brew install libpcap
     ```

3. **克隆项目**：
   ```bash
   git clone https://github.com/mikannse/goecap.git
   cd goecap