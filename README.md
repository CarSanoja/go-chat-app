# Chat Terminal TCP Application

This is a simple chat application built using Go. It supports multiple concurrent connections using goroutines and channels. The server and client communicate over TCP.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installing Dependencies](#installing-dependencies)
  - [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [Running the Client](#running-the-client)
- [License](#license)

## Getting Started

To get started with this project, you need to install the necessary dependencies and set up the configuration.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

### Installing Dependencies

Run the following command to install the necessary dependencies:

```
go mod tidy
```

Running the Server

To run the server, execute the following command:

```
go run cmd/server/main.go
```
The server will start and listen on the address and port specified in the config.yaml file (default is localhost:8080).
Running the Client

To run the client, execute the following command:

```
go run cmd/client/main.go
```

You will be prompted to enter your username and then you can start sending messages. Open multiple terminals to simulate multiple clients.