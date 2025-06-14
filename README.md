# VCS-ritz
## Ritz - A Git-Compatible Version Control System

Ritz is a lightweight, Git-style distributed version control system built in Go. It supports essential Git commands and internal mechanisms like Merkle tree-based commit history and zlib-compressed blob storage.

##  Features

- Git-compatible CLI structure (`ritz init`, `ritz add`, `ritz commit`, etc.)
- Merkle trees for immutable commit history
- Zlib compression for efficient file storage
- Basic branch management and checkout
- Custom packet-line protocol architecture

##  Installation

```bash
git clone https://github.com/your-username/ritz.git
cd ritz
go build -o my_cli .
