# Blink

Blink is a simple cross-platform notification app written in Go. It periodically sends desktop notifications to help you remember to blink, in order to help with dry eyes issues.
The idea came from reading this paper on the topic: https://pubmed.ncbi.nlm.nih.gov/40467388/

## Features

- Sends desktop notifications at a configurable interval (every 15 seconds by default, change with the `-interval` flag)
- Works on macOS (Intel & Apple Silicon), Linux, and Windows
- Lightweight and easy to use

## Installation

### Download Binaries

Go to the [Releases](https://github.com/yourusername/blink/releases) page and download the appropriate binary for your operating system:

- **macOS (Intel):** `blink-mac-amd64.tar.gz`
- **macOS (Apple Silicon):** `blink-mac-arm64.tar.gz`
- **Linux:** `blink-linux.tar.gz`
- **Windows:** `blink-win.zip`

Extract the archive and place the binary in your desired location.

### Build from Source

You need [Go installed](https://golang.org/doc/install).

```sh
git clone https://github.com/yourusername/blink.git
cd blink
go build -o blink
