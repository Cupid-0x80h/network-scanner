
# ⚡ Network Port Scanner (Go & Python)

Welcome to the Network Scanner project! This tool was built to scan open TCP ports on a given IP address using both Python and Go. It leverages the TCP three-way handshake process and uses concurrency for speed and performance — especially in the Go version.

---

## 📘 About

This project is both a **learning journal** and a **fully functioning port scanner**, created while studying how networks and TCP connections work.

The key concepts covered:
- What are ports?
- How does TCP handshaking work?
- How to implement a port scanner
- Performance issues and the need for concurrency
- How Go handles concurrency using Goroutines and Channels

---

## 🚀 Features

- Simple scanner using Python
- Basic and fast scanner in Go
- Concurrency model using Go routines
- Interactive CLI with user input
- Option to scan:
  - Default ports (1–1024)
  - All ports (1–65535)

---

## 🛠️ Usage

### ✅ Python Version (Basic)
```bash
python3 scanner.py
```
> This checks a single port on a single IP using sockets.

---

### 🦫 Go Version (Fast & Concurrent)

#### 1. Compile the scanner:
```bash
go build scanner.go
```

#### 2. Run the executable:
```bash
./scanner
```

#### 3. Follow the prompts:
- Enter target IP (e.g., `scanme.nmap.org`)
- Choose scan type (1 for ports 1-1024, 2 for full scan)

---

## 🧠 Conceptual Highlights

- **TCP Handshake**: Uses SYN and SYN/ACK to determine open ports.
- **Concurrency**: 100 worker goroutines scan ports in parallel.
- **Timeouts**: 0.5–1 second timeouts prevent hanging on filtered ports.
- **Channels**: Work queue (ports to scan) and results queue for open ports.

---

## 📂 Project Structure

```
.
├── scanner.go         # Final Go-based scanner with concurrency and menu
├── scanner.py         # Simple Python scanner
├── README.md          # You're reading it
└── network-scanner-docs.pdf  # Learning material and documentation
```

---

## 👤 Author

**Upesh Bhujel** aka `Cupid`  
[GitHub](https://github.com/Cupid-0x80h) • [LinkedIn](https://www.linkedin.com/in/bhujelupesh/)

> This project was a personal learning experience that evolved into a usable tool. Feel free to use or modify the code and docs.

---


