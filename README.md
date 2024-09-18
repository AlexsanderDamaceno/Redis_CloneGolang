
<div >
  <img src="https://upload.wikimedia.org/wikipedia/commons/6/64/Logo-redis.svg" alt="Redis Logo" margin="left" width="200" height="200">
</div>

# Redis Clone

## Clone
**Redis Clone** is a project implemented in ``Go``. It aims to ``replicate`` the ``basic functionality`` of ``Redis``.

## Features
* **Key-Value pair Store**: Store and get data using a  key-value pair interface.
* **Networking**: Communicate with ``Redis Server`` using TCP/IP protocol.

## Installation

⚠️ If you do not have **Go**, install using [Go Install Documentation](https://go.dev/doc/install)

Clone the repository:
```bash
$ git clone https://github.com/AlexsanderDamaceno/Redis_CloneGolang
```
Run the project:
```bash
$ go run main.go
```

### Usage:
Connect to the server using a Telnet:
```bash
$ telnet localhost 1234

# Try basic commands. For example, to test server commands:

$ set test_key test_value
$ get test_key 
