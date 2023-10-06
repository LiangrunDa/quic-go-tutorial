# quic-go Tutorial

This repo is for a TUM course project. The tutorial will walk you through the process of writing a simple QUIC client and server. It includes the following features:

- Basic quic client and server implementation for file transfer
- ZeroRTT handshake
- Chacha20 encryption
- Version negotiation
- In memory optimized server

At the time of this course project was done, zeroRTT is already implemented but not yet released. So we have to use the vendor mode of go module to use the latest version of [quic-go](https://github.com/quic-go/quic-go/). The vendor mode is enabled by default in this repo.

## Prerequisites

- Go 1.18 or higher

## Getting Started

```bash
make
```

You will find the binaries in the `testArtifact` directory including the client and server binaries. You can run the client and server in different terminals.